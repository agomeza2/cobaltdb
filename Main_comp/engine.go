package maincomp

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

//
// ----------- Estructuras -----------
//

type ObjectRef struct {
	User string `json:"user"`
	DB   string `json:"db"`
	Type string `json:"type"` // "node" o "relation"
	Name string `json:"name"`
	ID   int    `json:"id"`
}

type Engine struct {
	Users    sync.Map // map[string]*UserData
	Index    sync.Map // map[string]map[string][]ObjectRef
	basePath string
}

type UserData struct {
	Databases sync.Map // map[string]*Database
}

type Database struct {
	Nodes         sync.Map // map[string]*sync.Map[int]*Node
	Relations     sync.Map // map[string]*sync.Map[int]*Relation
	NodeRelations sync.Map // map[int]*sync.Map[int]*Relation
}

//
// ----------- Constructor -----------
//

func NewEngine(basePath string) *Engine {
	return &Engine{
		basePath: basePath,
	}
}

//
// ----------- Gestión de usuarios y DB -----------
//

func (eng *Engine) CreateUser(name string) error {
	_, loaded := eng.Users.LoadOrStore(name, &UserData{})
	if loaded {
		return fmt.Errorf("user %s already exists", name)
	}
	return nil
}

func (eng *Engine) DeleteUser(name string) error {
	eng.Users.Delete(name)
	return nil
}
func (eng *Engine) ListDatabases(user string) ([]string, error) {
	uRaw, ok := eng.Users.Load(user)
	if !ok {
		return nil, fmt.Errorf("user %s not found", user)
	}
	u := uRaw.(*UserData)

	var dbs []string
	u.Databases.Range(func(key, _ interface{}) bool {
		dbs = append(dbs, key.(string))
		return true
	})

	return dbs, nil
}

func (eng *Engine) CreateDatabase(user, dbName string) error {
	uRaw, ok := eng.Users.Load(user)
	if !ok {
		return fmt.Errorf("user %s not found", user)
	}
	u := uRaw.(*UserData)
	_, loaded := u.Databases.LoadOrStore(dbName, &Database{})
	if loaded {
		return fmt.Errorf("database %s already exists", dbName)
	}
	return nil
}

func (eng *Engine) DeleteDatabase(user, dbName string) error {
	uRaw, ok := eng.Users.Load(user)
	if !ok {
		return fmt.Errorf("user %s not found", user)
	}
	u := uRaw.(*UserData)
	u.Databases.Delete(dbName)
	return nil
}

//
// ----------- Creación de nodos y relaciones -----------
//

func (eng *Engine) CreateNode(user, dbName, name, category string, properties map[string]interface{}) (*Node, error) {
	uRaw, ok := eng.Users.Load(user)
	if !ok {
		return nil, fmt.Errorf("user %s not found", user)
	}
	u := uRaw.(*UserData)
	dbRaw, ok := u.Databases.Load(dbName)
	if !ok {
		return nil, fmt.Errorf("database %s not found", dbName)
	}
	db := dbRaw.(*Database)

	node := NewNode(category, name, properties)

	nameMapRaw, _ := db.Nodes.LoadOrStore(name, &sync.Map{})
	nameMap := nameMapRaw.(*sync.Map)
	nameMap.Store(node.ID, node)
	db.NodeRelations.Store(node.ID, &sync.Map{})

	// Indexar propiedades concurrente
	for k, v := range properties {
		go func(k string, v interface{}) {
			valStr := fmt.Sprintf("%v", v)
			fieldRaw, _ := eng.Index.LoadOrStore(k, &sync.Map{})
			field := fieldRaw.(*sync.Map)

			valRaw, _ := field.LoadOrStore(valStr, &sync.Map{})
			valMap := valRaw.(*sync.Map)
			valMap.Store(node.ID, ObjectRef{User: user, DB: dbName, Type: "node", Name: name, ID: node.ID})
		}(k, v)
	}

	return node, nil
}
// addNode agrega un nodo ya existente al Engine (sin crear uno nuevo).
func (eng *Engine) AddNode(user, dbName string, node *Node) error {
	uRaw, ok := eng.Users.Load(user)
	if !ok {
		return fmt.Errorf("user %s not found", user)
	}
	u := uRaw.(*UserData)

	dbRaw, ok := u.Databases.Load(dbName)
	if !ok {
		return fmt.Errorf("database %s not found", dbName)
	}
	db := dbRaw.(*Database)

	// Guardar el nodo en el mapa correspondiente
	nameMapRaw, _ := db.Nodes.LoadOrStore(node.Name, &sync.Map{})
	nameMap := nameMapRaw.(*sync.Map)
	nameMap.Store(node.ID, node)

	// Asegurar NodeRelations vacío
	db.NodeRelations.LoadOrStore(node.ID, &sync.Map{})

	// Actualizar índice
	for k, v := range node.Properties {
		valStr := fmt.Sprintf("%v", v)
		fieldRaw, _ := eng.Index.LoadOrStore(k, &sync.Map{})
		field := fieldRaw.(*sync.Map)

		valRaw, _ := field.LoadOrStore(valStr, &sync.Map{})
		valMap := valRaw.(*sync.Map)
		valMap.Store(node.ID, ObjectRef{
			User: user,
			DB:   dbName,
			Type: "node",
			Name: node.Name,
			ID:   node.ID,
		})
	}

	return nil
}
   
func (eng *Engine) CreateRelation(user, dbName, origin, destination, name, category string, properties map[string]interface{}) (*Relation, error) {
	uRaw, ok := eng.Users.Load(user)
	if !ok {
		return nil, fmt.Errorf("user %s not found", user)
	}
	u := uRaw.(*UserData)
	dbRaw, ok := u.Databases.Load(dbName)
	if !ok {
		return nil, fmt.Errorf("database %s not found", dbName)
	}
	db := dbRaw.(*Database)

	sourceNodesRaw, ok1 := db.Nodes.Load(origin)
	targetNodesRaw, ok2 := db.Nodes.Load(destination)
	if !ok1 || !ok2 {
		return nil, fmt.Errorf("source or target node not found")
	}

	var sourceNode, targetNode *Node
	sourceNodes := sourceNodesRaw.(*sync.Map)
	targetNodes := targetNodesRaw.(*sync.Map)
	sourceNodes.Range(func(_, v interface{}) bool {
		sourceNode = v.(*Node)
		return false
	})
	targetNodes.Range(func(_, v interface{}) bool {
		targetNode = v.(*Node)
		return false
	})

	relation := NewRelation(sourceNode, targetNode, category, name, properties)

	nameMapRaw, _ := db.Relations.LoadOrStore(name, &sync.Map{})
	nameMap := nameMapRaw.(*sync.Map)
	nameMap.Store(relation.ID, relation)

	// Actualizar NodeRelations
	updateNodeRelations := func(nodeID int) {
		nRelRaw, _ := db.NodeRelations.LoadOrStore(nodeID, &sync.Map{})
		nRel := nRelRaw.(*sync.Map)
		nRel.Store(relation.ID, relation)
	}
	updateNodeRelations(sourceNode.ID)
	updateNodeRelations(targetNode.ID)

	// Indexar propiedades concurrente
	for k, v := range properties {
		go func(k string, v interface{}) {
			valStr := fmt.Sprintf("%v", v)
			fieldRaw, _ := eng.Index.LoadOrStore(k, &sync.Map{})
			field := fieldRaw.(*sync.Map)
			valRaw, _ := field.LoadOrStore(valStr, &sync.Map{})
			valMap := valRaw.(*sync.Map)
			valMap.Store(relation.ID, ObjectRef{User: user, DB: dbName, Type: "relation", Name: name, ID: relation.ID})
		}(k, v)
	}

	return relation, nil
}
// AddRelation agrega una relación ya existente al Engine.
func (eng *Engine) AddRelation(user, dbName string, relation *Relation) error {
	uRaw, ok := eng.Users.Load(user)
	if !ok {
		return fmt.Errorf("user %s not found", user)
	}
	u := uRaw.(*UserData)

	dbRaw, ok := u.Databases.Load(dbName)
	if !ok {
		return fmt.Errorf("database %s not found", dbName)
	}
	db := dbRaw.(*Database)

	// Guardar relación en el mapa
	nameMapRaw, _ := db.Relations.LoadOrStore(relation.Name, &sync.Map{})
	nameMap := nameMapRaw.(*sync.Map)
	nameMap.Store(relation.ID, relation)

	// Actualizar NodeRelations (para origen y destino)
	updateNodeRelations := func(nodeID int) {
		nRelRaw, _ := db.NodeRelations.LoadOrStore(nodeID, &sync.Map{})
		nRel := nRelRaw.(*sync.Map)
		nRel.Store(relation.ID, relation)
	}
	updateNodeRelations(relation.Source.ID)
	updateNodeRelations(relation.Target.ID)

	// Actualizar índice
	for k, v := range relation.Properties {
		valStr := fmt.Sprintf("%v", v)
		fieldRaw, _ := eng.Index.LoadOrStore(k, &sync.Map{})
		field := fieldRaw.(*sync.Map)
		valRaw, _ := field.LoadOrStore(valStr, &sync.Map{})
		valMap := valRaw.(*sync.Map)
		valMap.Store(relation.ID, ObjectRef{
			User: user,
			DB:   dbName,
			Type: "relation",
			Name: relation.Name,
			ID:   relation.ID,
		})
	}

	return nil
}

//
// ----------- Borrado -----------
//

func (eng *Engine) DeleteNode(user, dbName, name string, id int) error {
	uRaw, ok := eng.Users.Load(user)
	if !ok {
		return fmt.Errorf("user %s not found", user)
	}
	u := uRaw.(*UserData)
	dbRaw, ok := u.Databases.Load(dbName)
	if !ok {
		return fmt.Errorf("database %s not found", dbName)
	}
	db := dbRaw.(*Database)

	nodesRaw, ok := db.Nodes.Load(name)
	if !ok {
		return fmt.Errorf("node name not found")
	}
	nodes := nodesRaw.(*sync.Map)
	nodes.Delete(id)
	db.NodeRelations.Delete(id)
	return nil
}

func (eng *Engine) DeleteRelation(user, dbName, name string, id int) error {
	uRaw, ok := eng.Users.Load(user)
	if !ok {
		return fmt.Errorf("user %s not found", user)
	}
	u := uRaw.(*UserData)
	dbRaw, ok := u.Databases.Load(dbName)
	if !ok {
		return fmt.Errorf("database %s not found", dbName)
	}
	db := dbRaw.(*Database)

	relsRaw, ok := db.Relations.Load(name)
	if !ok {
		return nil
	}
	rels := relsRaw.(*sync.Map)
	relRaw, ok := rels.Load(id)
	if ok {
		rel := relRaw.(*Relation)
		db.NodeRelations.Delete(rel.Source.ID)
		db.NodeRelations.Delete(rel.Target.ID)
	}
	rels.Delete(id)
	return nil
}

//
// ----------- Búsqueda y resolución -----------
//

func (eng *Engine) FindByProperty(field, value string) []ObjectRef {
	fieldRaw, ok := eng.Index.Load(field)
	if !ok {
		return nil
	}
	fieldMap := fieldRaw.(*sync.Map)
	valRaw, ok := fieldMap.Load(value)
	if !ok {
		return nil
	}
	valMap := valRaw.(*sync.Map)
	var results []ObjectRef
	valMap.Range(func(_, v interface{}) bool {
		results = append(results, v.(ObjectRef))
		return true
	})
	return results
}

func (eng *Engine) Resolve(ref ObjectRef) (interface{}, error) {
	uRaw, ok := eng.Users.Load(ref.User)
	if !ok {
		return nil, fmt.Errorf("user not found")
	}
	u := uRaw.(*UserData)
	dbRaw, ok := u.Databases.Load(ref.DB)
	if !ok {
		return nil, fmt.Errorf("database not found")
	}
	db := dbRaw.(*Database)

	switch ref.Type {
	case "node":
		nodesRaw, ok := db.Nodes.Load(ref.Name)
		if !ok {
			return nil, fmt.Errorf("node not found")
		}
		nodes := nodesRaw.(*sync.Map)
		nRaw, ok := nodes.Load(ref.ID)
		if !ok {
			return nil, fmt.Errorf("node not found")
		}
		return nRaw, nil
	case "relation":
		relsRaw, ok := db.Relations.Load(ref.Name)
		if !ok {
			return nil, fmt.Errorf("relation not found")
		}
		rels := relsRaw.(*sync.Map)
		rRaw, ok := rels.Load(ref.ID)
		if !ok {
			return nil, fmt.Errorf("relation not found")
		}
		return rRaw, nil
	}
	return nil, fmt.Errorf("unknown type")
}

func (eng *Engine) RelationsOfNode(user, dbName string, nodeID int) []*Relation {
	uRaw, ok := eng.Users.Load(user)
	if !ok {
		return nil
	}
	u := uRaw.(*UserData)
	dbRaw, ok := u.Databases.Load(dbName)
	if !ok {
		return nil
	}
	db := dbRaw.(*Database)

	nRelRaw, ok := db.NodeRelations.Load(nodeID)
	if !ok {
		return nil
	}
	nRel := nRelRaw.(*sync.Map)
	var result []*Relation
	nRel.Range(func(_, v interface{}) bool {
		result = append(result, v.(*Relation))
		return true
	})
	return result
}

// ----------- Persistencia concurrente -----------
func (eng *Engine) LoadDB(user, dbName string) error {
	uRaw, ok := eng.Users.Load(user)
	if !ok {
		return fmt.Errorf("user %s not found", user)
	}
	u := uRaw.(*UserData)

	db := &Database{}
	u.Databases.Store(dbName, db)

	path := filepath.Join(eng.basePath, user, dbName)

	// Cargar nodos
	nodesFile := filepath.Join(path, "nodes.json")
	nodesData := make(map[string]map[int]*Node)
	if _, err := os.Stat(nodesFile); err == nil {
		data, err := os.ReadFile(nodesFile)
		if err != nil {
			return fmt.Errorf("failed to read nodes.json: %v", err)
		}
		if err := json.Unmarshal(data, &nodesData); err != nil {
			return fmt.Errorf("failed to unmarshal nodes.json: %v", err)
		}
		for name, idMap := range nodesData {
			nodeMap := &sync.Map{}
			for id, node := range idMap {
				nodeMap.Store(id, node)
				db.NodeRelations.Store(id, &sync.Map{})
			}
			db.Nodes.Store(name, nodeMap)
		}
	}

	// Cargar relaciones
	relsFile := filepath.Join(path, "relations.json")
	relsData := make(map[string]map[int]*Relation)
	if _, err := os.Stat(relsFile); err == nil {
		data, err := os.ReadFile(relsFile)
		if err != nil {
			return fmt.Errorf("failed to read relations.json: %v", err)
		}
		if err := json.Unmarshal(data, &relsData); err != nil {
			return fmt.Errorf("failed to unmarshal relations.json: %v", err)
		}
		for name, idMap := range relsData {
			relMap := &sync.Map{}
			for id, rel := range idMap {
				relMap.Store(id, rel)
				// Reconstruir NodeRelations
				nRelRaw, _ := db.NodeRelations.LoadOrStore(rel.Source.ID, &sync.Map{})
				nRel := nRelRaw.(*sync.Map)
				nRel.Store(rel.ID, rel)

				nRelRaw2, _ := db.NodeRelations.LoadOrStore(rel.Target.ID, &sync.Map{})
				nRel2 := nRelRaw2.(*sync.Map)
				nRel2.Store(rel.ID, rel)

				// Reconstruir índice
				for k, v := range rel.Properties {
					valStr := fmt.Sprintf("%v", v)
					fieldRaw, _ := eng.Index.LoadOrStore(k, &sync.Map{})
					field := fieldRaw.(*sync.Map)
					valRaw, _ := field.LoadOrStore(valStr, &sync.Map{})
					valMap := valRaw.(*sync.Map)
					valMap.Store(rel.ID, ObjectRef{User: user, DB: dbName, Type: "relation", Name: name, ID: id})
				}
			}
			db.Relations.Store(name, relMap)
		}
	}

	// Reconstruir índice de nodos
	db.Nodes.Range(func(name, nMapRaw interface{}) bool {
		nMap := nMapRaw.(*sync.Map)
		nMap.Range(func(id, nodeRaw interface{}) bool {
			node := nodeRaw.(*Node)
			for k, v := range node.Properties {
				valStr := fmt.Sprintf("%v", v)
				fieldRaw, _ := eng.Index.LoadOrStore(k, &sync.Map{})
				field := fieldRaw.(*sync.Map)
				valRaw, _ := field.LoadOrStore(valStr, &sync.Map{})
				valMap := valRaw.(*sync.Map)
				valMap.Store(node.ID, ObjectRef{User: user, DB: dbName, Type: "node", Name: name.(string), ID: node.ID})
			}
			return true
		})
		return true
	})

	return nil
}

func (eng *Engine) SaveDB(user, dbName string) error {
	uRaw, ok := eng.Users.Load(user)
	if !ok {
		return fmt.Errorf("user not found")
	}
	u := uRaw.(*UserData)
	dbRaw, ok := u.Databases.Load(dbName)
	if !ok {
		return fmt.Errorf("database not found")
	}
	db := dbRaw.(*Database)

	path := filepath.Join(eng.basePath, user, dbName)
	os.MkdirAll(path, 0755)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		nodesData := make(map[string]map[int]*Node)
		db.Nodes.Range(func(k, v interface{}) bool {
			idMap := make(map[int]*Node)
			v.(*sync.Map).Range(func(id, n interface{}) bool {
				idMap[id.(int)] = n.(*Node)
				return true
			})
			nodesData[k.(string)] = idMap
			return true
		})
		data, _ := json.Marshal(nodesData)
		os.WriteFile(filepath.Join(path, "nodes.json"), data, 0644)
	}()

	go func() {
		defer wg.Done()
		relsData := make(map[string]map[int]*Relation)
		db.Relations.Range(func(k, v interface{}) bool {
			idMap := make(map[int]*Relation)
			v.(*sync.Map).Range(func(id, r interface{}) bool {
				idMap[id.(int)] = r.(*Relation)
				return true
			})
			relsData[k.(string)] = idMap
			return true
		})
		data, _ := json.Marshal(relsData)
		os.WriteFile(filepath.Join(path, "relations.json"), data, 0644)
	}()

	wg.Wait()
	return nil
}
