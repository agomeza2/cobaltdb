package l27

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sort"
	"strings"
)

// Estructuras internas
type GraphDB struct {
	Databases map[string]*DBInstance
	Current   string
	Users     map[string]string // username -> role
}

type DBInstance struct {
	NodeCounter int
	RelCounter  int
	Nodes       []map[string]interface{}
	Relations   []map[string]interface{}
}

func NewGraphDB() *GraphDB {
	return &GraphDB{
		Databases: map[string]*DBInstance{},
		Users:     map[string]string{},
	}
}

func NewDBInstance() *DBInstance {
	return &DBInstance{
		NodeCounter: 0,
		RelCounter:  0,
		Nodes:       []map[string]interface{}{},
		Relations:   []map[string]interface{}{},
	}
}

/* ----------------- Dispatcher ----------------- */
func (g *GraphDB) Execute(st *Stmt) {
	switch {
	case st.Create != nil:
		g.execCreate(st.Create)
	case st.CreateUser != nil:
		g.execCreateUser(st.CreateUser)
	case st.CreateDB != nil:
		g.execCreateDB(st.CreateDB)
	case st.RemoveDB != nil:
		g.execRemoveDB(st.RemoveDB)
	case st.Use != nil:
		g.execUse(st.Use)
	case st.Show != nil:
		g.execShow(st.Show)
	case st.List != nil:
		g.execList(st.List)
	case st.Alter != nil:
		g.execAlter(st.Alter)
	case st.RemoveAttr != nil:
		g.execRemoveAttr(st.RemoveAttr)
	case st.AddAttr != nil:
		g.execAddAttr(st.AddAttr)
	case st.HowMany != nil:
		g.execHowMany(st.HowMany)
	case st.Common != nil:
		g.execCommon(st.Common)
	case st.Search != nil:
		g.execSearch(st.Search)
	case st.Import != nil:
		g.execImport(st.Import)
	case st.Export != nil:
		g.execExport(st.Export)
	default:
		fmt.Println("Comando no reconocido.")
	}
}

/* ----------------- CREATE family ----------------- */
func (g *GraphDB) execCreate(c *CreateStmt) {
	switch strings.ToUpper(c.What) {
	case "DATABASE":
		if c.DBName == nil {
			fmt.Println("CREATE DATABASE requiere un nombre.")
			return
		}
		name := *c.DBName
		if _, ok := g.Databases[name]; ok {
			fmt.Println("DATABASE ya existe:", name)
			return
		}
		g.Databases[name] = NewDBInstance()
		fmt.Println("DATABASE creada:", name)
	case "NODE":
		inst := g.getCurrent()
		if inst == nil {
			fmt.Println("No hay database seleccionada (USE <name>).")
			return
		}
		if c.Obj == nil {
			fmt.Println("CREATE NODE requiere un objeto JSON.")
			return
		}
		m := objToMap(c.Obj)
		inst.NodeCounter++
		m["id"] = inst.NodeCounter
		inst.Nodes = append(inst.Nodes, m)
		fmt.Println("NODE creado id:", inst.NodeCounter)
	case "NODES":
		inst := g.getCurrent()
		if inst == nil {
			fmt.Println("No hay database seleccionada.")
			return
		}
		if len(c.ArrayObjs) == 0 {
			fmt.Println("CREATE NODES requiere una lista de objetos.")
			return
		}
		for _, o := range c.ArrayObjs {
			m := objToMap(o)
			inst.NodeCounter++
			m["id"] = inst.NodeCounter
			inst.Nodes = append(inst.Nodes, m)
			fmt.Println("NODE creado id:", inst.NodeCounter)
		}
	case "RELATION":
		inst := g.getCurrent()
		if inst == nil {
			fmt.Println("No hay database seleccionada.")
			return
		}
		if c.Obj == nil {
			fmt.Println("CREATE RELATION requiere un objeto JSON.")
			return
		}
		m := objToMap(c.Obj)
		inst.RelCounter++
		m["id"] = inst.RelCounter
		inst.Relations = append(inst.Relations, m)
		fmt.Println("RELATION creada id:", inst.RelCounter)
	case "RELATIONS":
		inst := g.getCurrent()
		if inst == nil {
			fmt.Println("No hay database seleccionada.")
			return
		}
		for _, o := range c.ArrayObjs {
			m := objToMap(o)
			inst.RelCounter++
			m["id"] = inst.RelCounter
			inst.Relations = append(inst.Relations, m)
			fmt.Println("RELATION creada id:", inst.RelCounter)
		}
	default:
		fmt.Println("CREATE:", c.What, "no soportado.")
	}
}

func (g *GraphDB) execCreateUser(u *CreateUserStmt) {
	name := u.Name
	role := strings.ToUpper(u.Role)
	g.Users[name] = role
	fmt.Printf("Usuario creado: %s (%s)\n", name, role)
}

/* ----------------- DB control ----------------- */
func (g *GraphDB) execCreateDB(c *CreateDatabaseStmt) {
	name := c.Name
	if _, ok := g.Databases[name]; ok {
		fmt.Println("DATABASE ya existe:", name)
		return
	}
	g.Databases[name] = NewDBInstance()
	fmt.Println("DATABASE creada:", name)
}

func (g *GraphDB) execRemoveDB(r *RemoveDatabaseStmt) {
	name := r.Name
	if _, ok := g.Databases[name]; !ok {
		fmt.Println("DATABASE no existe:", name)
		return
	}
	delete(g.Databases, name)
	if g.Current == name {
		g.Current = ""
	}
	fmt.Println("DATABASE eliminada:", name)
}

func (g *GraphDB) execUse(u *UseStmt) {
	name := u.Name
	if _, ok := g.Databases[name]; !ok {
		fmt.Println("DATABASE no existe:", name)
		return
	}
	g.Current = name
	fmt.Println("Usando DATABASE:", name)
}

/* ----------------- SHOW / LIST ----------------- */
func (g *GraphDB) execShow(s *ShowStmt) {
	switch strings.ToUpper(s.Type) {
	case "DATABASES":
		for name := range g.Databases {
			fmt.Println(name)
		}
	case "NODES", "NODE":
		inst := g.getCurrent()
		if inst == nil {
			fmt.Println("No hay database seleccionada.")
			return
		}
		if s.Filter == nil {
			for _, n := range inst.Nodes {
				fmt.Println(prettyJSON(n))
			}
			return
		}
		for _, n := range inst.Nodes {
			if matchFilter(n, s.Filter) {
				fmt.Println(prettyJSON(n))
			}
		}
	case "RELATIONS", "RELATION":
		inst := g.getCurrent()
		if inst == nil {
			fmt.Println("No hay database seleccionada.")
			return
		}
		if s.Filter == nil {
			for _, r := range inst.Relations {
				fmt.Println(prettyJSON(r))
			}
			return
		}
		for _, r := range inst.Relations {
			if matchFilter(r, s.Filter) {
				fmt.Println(prettyJSON(r))
			}
		}
	default:
		fmt.Println("SHOW no soporta:", s.Type)
	}
}

func (g *GraphDB) execList(l *ListStmt) {
	switch strings.ToUpper(l.What) {
	case "DATABASES":
		for name := range g.Databases {
			fmt.Println(name)
		}
	case "NODES":
		inst := g.getCurrent()
		if inst == nil {
			fmt.Println("No hay database seleccionada.")
			return
		}
		for _, n := range inst.Nodes {
			fmt.Println(prettyJSON(n))
		}
	case "RELATIONS":
		inst := g.getCurrent()
		if inst == nil {
			fmt.Println("No hay database seleccionada.")
			return
		}
		for _, r := range inst.Relations {
			fmt.Println(prettyJSON(r))
		}
	default:
		fmt.Println("LIST no soporta:", l.What)
	}
}

/* ----------------- ALTER / ADD / REMOVE ATTR ----------------- */
func (g *GraphDB) execAlter(a *AlterStmt) {
	if a == nil || a.Object == nil {
		fmt.Println("ALTER requiere objetivo y cuerpo.")
		return
	}
	inst := g.getCurrent()
	if inst == nil {
		fmt.Println("No hay database seleccionada.")
		return
	}
	target := a.Object
	// buscar nodos/relations que cumplan target.Filter
	switch strings.ToUpper(allowString(target.What)) {
	case "NODE", "":
		for _, n := range inst.Nodes {
			if matchFilter(n, target.Filter) {
				if a.SetBody != nil {
					mergeMap(n, objToMap(a.SetBody))
					fmt.Printf("Nodo id:%v alterado\n", n["id"])
				}
			}
		}
	case "RELATION":
		for _, r := range inst.Relations {
			if matchFilter(r, target.Filter) {
				if a.SetBody != nil {
					mergeMap(r, objToMap(a.SetBody))
					fmt.Printf("Relación id:%v alterada\n", r["id"])
				}
			}
		}
	default:
		fmt.Println("ALTER target no reconocido:", target.What)
	}
}

func (g *GraphDB) execRemoveAttr(r *RemoveAttrStmt) {
	inst := g.getCurrent()
	if inst == nil {
		fmt.Println("No hay database seleccionada.")
		return
	}
	switch strings.ToUpper(r.What) {
	case "NODE":
		for _, n := range inst.Nodes {
			if matchFilter(n, r.Filter) {
				delete(n, r.Key)
				fmt.Printf("Atributo %s eliminado en nodo id:%v\n", r.Key, n["id"])
			}
		}
	case "RELATION":
		for _, rel := range inst.Relations {
			if matchFilter(rel, r.Filter) {
				delete(rel, r.Key)
				fmt.Printf("Atributo %s eliminado en relación id:%v\n", r.Key, rel["id"])
			}
		}
	default:
		fmt.Println("REMOVE IN target no soportado:", r.What)
	}
}

func (g *GraphDB) execAddAttr(a *AddAttrStmt) {
	inst := g.getCurrent()
	if inst == nil {
		fmt.Println("No hay database seleccionada.")
		return
	}
	switch strings.ToUpper(a.What) {
	case "NODE":
		for _, n := range inst.Nodes {
			if matchFilter(n, a.Filter) {
				n[a.Key] = nil
				fmt.Printf("Atributo %s agregado (nil) en nodo id:%v\n", a.Key, n["id"])
			}
		}
	case "RELATION":
		for _, r := range inst.Relations {
			if matchFilter(r, a.Filter) {
				r[a.Key] = nil
				fmt.Printf("Atributo %s agregado (nil) en relación id:%v\n", a.Key, r["id"])
			}
		}
	default:
		fmt.Println("ADD IN target no soportado:", a.What)
	}
}

/* ----------------- HOW MANY ----------------- */
func (g *GraphDB) execHowMany(h *HowManyStmt) {
	inst := g.getCurrent()
	if inst == nil {
		fmt.Println("No hay database seleccionada.")
		return
	}
	search := strings.ToLower(h.Ident)
	countNodes := 0
	countRels := 0
	for _, n := range inst.Nodes {
		if v, ok := n["category"].(string); ok && strings.ToLower(v) == search {
			countNodes++
		}
	}
	for _, r := range inst.Relations {
		if v, ok := r["category"].(string); ok && strings.ToLower(v) == search {
			countRels++
		}
	}
	fmt.Printf("HOW MANY %s => nodes:%d relations:%d\n", h.Ident, countNodes, countRels)
}

/* ----------------- COMMON ----------------- */
func (g *GraphDB) execCommon(c *CommonStmt) {
	if c == nil || len(c.Filters) < 2 {
		fmt.Println("COMMON requiere 2 o más filtros.")
		return
	}
	inst := g.getCurrent()
	if inst == nil {
		fmt.Println("No hay database seleccionada.")
		return
	}
	sets := []map[interface{}]bool{}
	for _, f := range c.Filters {
		set := map[interface{}]bool{}
		for _, n := range inst.Nodes {
			if matchFilter(n, f) {
				set[n["id"]] = true
			}
		}
		sets = append(sets, set)
	}
	// intersección
	res := map[interface{}]bool{}
	for k := range sets[0] {
		ok := true
		for i := 1; i < len(sets); i++ {
			if !sets[i][k] {
				ok = false
				break
			}
		}
		if ok {
			res[k] = true
		}
	}
	if len(res) == 0 {
		fmt.Println("COMMON: sin resultados.")
		return
	}
	for id := range res {
		if n := findNodeByID(g.getCurrent(), id); n != nil {
			fmt.Println(prettyJSON(n))
		}
	}
}

/* ----------------- SEARCH (WHERE / GROUP BY / ORDER) ----------------- */
func (g *GraphDB) execSearch(s *SearchStmt) {
	inst := g.getCurrent()
	if inst == nil {
		fmt.Println("No hay database seleccionada.")
		return
	}
	// filtrar nodes por tipo (Type puede referirse a category)
	results := []map[string]interface{}{}
	for _, n := range inst.Nodes {
		// si s.Type es una categoría, filtrar; si no, dejar
		if s.Type != "" {
			if cat, ok := n["category"].(string); ok && strings.ToLower(cat) != strings.ToLower(s.Type) {
				continue
			}
		}
		if s.Where != nil {
			if !matchExpression(n, s.Where.Expr) {
				continue
			}
		}
		results = append(results, n)
	}
	// group by (simple: print grouped buckets)
	if s.GroupBy != nil {
		groups := map[string][]map[string]interface{}{}
		for _, r := range results {
			key := fmt.Sprintf("%v", getNestedValue(r, s.GroupBy.Field))
			groups[key] = append(groups[key], r)
		}
		for k, group := range groups {
			fmt.Printf("GROUP: %s (%d)\n", k, len(group))
			for _, item := range group {
				fmt.Println(prettyJSON(item))
			}
		}
		return
	}
	// order by (simple)
	if s.OrderBy != nil {
		field := s.OrderBy.Field
		dir := "ASC"
		if s.OrderBy.Direction != nil {
			dir = *s.OrderBy.Direction
		}
		sort.Slice(results, func(i, j int) bool {
			vi := fmt.Sprintf("%v", getNestedValue(results[i], field))
			vj := fmt.Sprintf("%v", getNestedValue(results[j], field))
			if dir == "DESC" {
				return vi > vj
			}
			return vi < vj
		})
	}
	for _, r := range results {
		fmt.Println(prettyJSON(r))
	}
}

/* ----------------- IMPORT / EXPORT ----------------- */
func (g *GraphDB) execImport(im *ImportStmt) {
	if im == nil || im.Path == nil {
		fmt.Println("IMPORT requiere ruta entre comillas.")
		return
	}
	inst := g.getCurrent()
	if inst == nil {
		fmt.Println("No hay database seleccionada.")
		return
	}
	path := unquote(*im.Path)
	abs, _ := filepath.Abs(path)
	b, err := ioutil.ReadFile(abs)
	if err != nil {
		fmt.Println("Error leyendo archivo:", err)
		return
	}
	var arr []map[string]interface{}
	if err := json.Unmarshal(b, &arr); err != nil {
		fmt.Println("IMPORT: JSON inválido:", err)
		return
	}
	for _, o := range arr {
		if _, hasSource := o["source"]; hasSource {
			inst.RelCounter++
			o["id"] = inst.RelCounter
			inst.Relations = append(inst.Relations, o)
			fmt.Println("Relación importada id:", inst.RelCounter)
		} else {
			inst.NodeCounter++
			o["id"] = inst.NodeCounter
			inst.Nodes = append(inst.Nodes, o)
			fmt.Println("Nodo importado id:", inst.NodeCounter)
		}
	}
}

func (g *GraphDB) execExport(e *ExportStmt) {
	if e == nil {
		return
	}
	inst := g.getCurrent()
	if inst == nil {
		fmt.Println("No hay database seleccionada.")
		return
	}
	if e.Name != g.Current {
		fmt.Println("EXPORT solo funciona para la database seleccionada:", g.Current)
		return
	}
	out := map[string]interface{}{
		"nodes":     inst.Nodes,
		"relations": inst.Relations,
	}
	data, _ := json.MarshalIndent(out, "", "  ")
	if e.ToTok == nil {
		fmt.Println(string(data))
		return
	}
	path := unquote(*e.ToTok)
	abs, _ := filepath.Abs(path)
	if err := ioutil.WriteFile(abs, data, 0644); err != nil {
		fmt.Println("Error exportando:", err)
		return
	}
	fmt.Println("EXPORT completado a", abs)
}

/* ----------------- Helpers ----------------- */

func (g *GraphDB) getCurrent() *DBInstance {
	if g.Current == "" {
		return nil
	}
	return g.Databases[g.Current]
}

func objToMap(obj *JSONObject) map[string]interface{} {
	m := map[string]interface{}{}
	if obj == nil {
		return m
	}
	for _, kv := range obj.Pairs {
		m[kv.Key] = litToValue(kv.Value)
	}
	return m
}

func litToValue(l *Lit) interface{} {
	if l == nil {
		return nil
	}
	if l.Number != nil {
		return *l.Number
	}
	if l.Str != nil {
		return unquote(*l.Str)
	}
	if l.Ident != nil {
		return *l.Ident
	}
	return nil
}

func unquote(s string) string {
	if len(s) >= 2 && s[0] == '"' && s[len(s)-1] == '"' {
		return s[1 : len(s)-1]
	}
	return s
}

func mergeMap(dst, src map[string]interface{}) {
	for k, v := range src {
		dst[k] = v
	}
}

func allowString(p *string) string {
	if p == nil {
		return ""
	}
	return *p
}

func matchFilter(doc map[string]interface{}, f *Filter) bool {
	if f == nil || doc == nil {
		return false
	}
	key := f.Key
	val := litToValue(f.Value)
	// handle id specially
	if strings.ToLower(key) == "id" {
		return compareValues(doc["id"], val)
	}
	// try direct field
	if dv, ok := getNestedValueRaw(doc, key); ok {
		return compareValues(dv, val)
	}
	// fallback false
	return false
}

func matchExpression(doc map[string]interface{}, e *Expression) bool {
	if e == nil {
		return false
	}
	left := getNestedValue(doc, e.Left)
	right := e.Right
	if right == nil {
		return false
	}
	rVal := litToValue(right)
	switch e.Op {
	case "=":
		return compareValues(left, rVal)
	case "!=":
		return !compareValues(left, rVal)
	case ">":
		return compareNumeric(left, rVal, ">")
	case "<":
		return compareNumeric(left, rVal, "<")
	case ">=":
		return compareNumeric(left, rVal, ">=")
	case "<=":
		return compareNumeric(left, rVal, "<=")
	default:
		return false
	}
}

// obtiene valor de campo "a.b.c" (soporte básico)
func getNestedValue(doc map[string]interface{}, path string) interface{} {
	parts := strings.Split(path, ".")
	var cur interface{} = doc
	for _, p := range parts {
		if m, ok := cur.(map[string]interface{}); ok {
			cur = m[p]
		} else {
			return nil
		}
	}
	return cur
}

func getNestedValueRaw(doc map[string]interface{}, path string) (interface{}, bool) {
	v := getNestedValue(doc, path)
	if v == nil {
		return nil, false
	}
	return v, true
}

func compareValues(a, b interface{}) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	// numeric checks
	switch av := a.(type) {
	case int:
		switch bv := b.(type) {
		case int:
			return av == bv
		case string:
			return fmt.Sprintf("%d", av) == bv
		case float64:
			return float64(av) == bv
		}
	case float64:
		switch bv := b.(type) {
		case float64:
			return av == bv
		case int:
			return av == float64(bv)
		case string:
			return fmt.Sprintf("%g", av) == bv
		}
	case string:
		switch bv := b.(type) {
		case string:
			return av == bv
		default:
			return av == fmt.Sprintf("%v", bv)
		}
	}
	// fallback to string compare
	return fmt.Sprintf("%v", a) == fmt.Sprintf("%v", b)
}

func compareNumeric(a, b interface{}, op string) bool {
	var af, bf float64
	switch v := a.(type) {
	case int:
		af = float64(v)
	case float64:
		af = v
	case string:
		fmt.Sscanf(v, "%f", &af)
	default:
		return false
	}
	switch v := b.(type) {
	case int:
		bf = float64(v)
	case float64:
		bf = v
	case string:
		fmt.Sscanf(v, "%f", &bf)
	default:
		return false
	}
	switch op {
	case ">":
		return af > bf
	case "<":
		return af < bf
	case ">=":
		return af >= bf
	case "<=":
		return af <= bf
	default:
		return false
	}
}

func findNodeByID(inst *DBInstance, id interface{}) map[string]interface{} {
	for _, n := range inst.Nodes {
		if compareValues(n["id"], id) {
			return n
		}
	}
	return nil
}

func prettyJSON(v interface{}) string {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return fmt.Sprintf("%v", v)
	}
	return string(b)
}
