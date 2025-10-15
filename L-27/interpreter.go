package l27

import (
<<<<<<< HEAD
	storage "cobaltdb-local/Storage"
=======
	engine "cobaltdb-local/Main_comp"
>>>>>>> d4edcdab84f14de434df0b3539422c949ad05a23
	"fmt"
)

type Interpreter struct {
	DBPath      string
	CurrentDB   string
	CurrentUser string
<<<<<<< HEAD
	idx         *storage.Index
=======
	eng         *engine.Engine
>>>>>>> d4edcdab84f14de434df0b3539422c949ad05a23
}

func NewInterpreter(dbpath string) (*Interpreter, error) {
	interp := &Interpreter{
		DBPath: dbpath,
<<<<<<< HEAD
		idx:    storage.NewIndex(dbpath),
	}

	// Uncomment and implement loading logic if needed
	// err := interp.idx.LoadFromDisk()
=======
		eng:    engine.NewEngine(dbpath),
	}

	// Uncomment and implement loading logic if needed
	// err := interp.eng.LoadFromDisk()
>>>>>>> d4edcdab84f14de434df0b3539422c949ad05a23
	// if err != nil {
	// 	return nil, err
	// }

	return interp, nil
}

/*func (i *Interpreter) Save() {
<<<<<<< HEAD
	i.idx.FlushToDisk()
=======
	i.eng.FlushToDisk()
>>>>>>> d4edcdab84f14de434df0b3539422c949ad05a23
}*/

func (i *Interpreter) Execute(cmd *Command) error {
	switch cmd.Name {
	case "ls":
		return i.cmdLs(cmd.Args)
	case "cd":
		return i.cmdSelect(cmd.Args)
	case "touch":
		return i.cmdCreate(cmd.Args)
	case "insert":
		return i.cmdInsert(cmd.Properties, cmd.Filters, cmd.Args)
	case "modify":
		return i.cmdModify(cmd.Properties, cmd.Filters, cmd.Args)
	case "rm":
		return i.cmdDelete(cmd.Args)
	case "grep":
		return i.cmdFind(cmd.RawQuery, cmd.Args)
	case "import":
		return i.cmdImport(cmd.Args)
	case "export":
		return i.cmdExport(cmd.Args)
	default:
		return fmt.Errorf("comando no implementado: %s", cmd.Name)
	}
}

func (i *Interpreter) loadDb(user, dBName string) error {
<<<<<<< HEAD
	i.idx.LoadDB(user, dBName)
=======
	i.eng.LoadDB(user, dBName)
>>>>>>> d4edcdab84f14de434df0b3539422c949ad05a23
	return nil
}
func (i *Interpreter) cmdLs(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("list requiere argumento")
	}

	return nil
}

func (i *Interpreter) cmdSelect(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("select requiere argumento")
	}
	/*
		switch args[0] {
		case "db":
			if len(args) < 2 {
				return fmt.Errorf("select db requiere nombre de base de datos")
			}
			dbName := args[1]
<<<<<<< HEAD
			dbs := i.idx.ListDatabases()
=======
			dbs := i.eng.ListDatabases()
>>>>>>> d4edcdab84f14de434df0b3539422c949ad05a23
			found := false
			for _, d := range dbs {
				if d == dbName {
					found = true
					break
				}
			}
			if !found {
				return fmt.Errorf("base de datos %s no encontrada", dbName)
			}
			i.CurrentDB = dbName
			i.CurrentColl = ""
			fmt.Println("Base de datos seleccionada:", i.CurrentDB)
			return nil

		case "collection":
			if i.CurrentDB == "" {
				return fmt.Errorf("no hay base de datos seleccionada")
			}
			if len(args) < 2 {
				return fmt.Errorf("select collection requiere nombre de colección")
			}
			colName := args[1]
<<<<<<< HEAD
			cols, err := i.idx.ListCollections(i.CurrentDB)
=======
			cols, err := i.eng.ListCollections(i.CurrentDB)
>>>>>>> d4edcdab84f14de434df0b3539422c949ad05a23
			if err != nil {
				return err
			}
			found := false
			for _, c := range cols {
				if c == colName {
					found = true
					break
				}
			}
			if !found {
				return fmt.Errorf("colección %s no encontrada en base de datos %s", colName, i.CurrentDB)
			}
			i.CurrentColl = colName
			fmt.Println("Colección seleccionada:", i.CurrentColl)
			return nil

		default:
			return fmt.Errorf("argumento desconocido para select: %s", args[0])
		}*/
	println("Comando select no implementado aún")
	return nil
}

// Implementa el resto con la lógica que necesites
func (i *Interpreter) cmdCreate(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("create needs an argument")
	} /*
		switch args[0] {
		case "db":
<<<<<<< HEAD
			i.idx.CreateDatabase(args[1])
		case "collections":
			i.idx.CreateCollection(i.CurrentDB, args[1])
		case "documents":
			i.idx.CreateDocument(i.CurrentDB, i.CurrentColl, args[1])
=======
			i.eng.CreateDatabase(args[1])
		case "collections":
			i.eng.CreateCollection(i.CurrentDB, args[1])
		case "documents":
			i.eng.CreateDocument(i.CurrentDB, i.CurrentColl, args[1])
>>>>>>> d4edcdab84f14de434df0b3539422c949ad05a23
		default:
			return fmt.Errorf("unkown argument for create: %s", args[0])
		}*/
	return nil
}

func (i *Interpreter) cmdInsert(props, filters []map[string]interface{}, args []string) error {
	fmt.Println("Comando insert no implementado aún")
	return nil
}

func (i *Interpreter) cmdModify(props, filters []map[string]interface{}, args []string) error {
	return nil
}

func (i *Interpreter) cmdDelete(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("delete needs an argument")
	}
	/*
		switch args[0] {
		case "db":
<<<<<<< HEAD
			i.idx.DeleteDatabase(args[1])
		case "collections":
			i.idx.DeleteCollection(i.CurrentDB, args[1])
		case "documents":
			i.idx.DeleteDocument(i.CurrentDB, i.CurrentColl, args[1])
=======
			i.eng.DeleteDatabase(args[1])
		case "collections":
			i.eng.DeleteCollection(i.CurrentDB, args[1])
		case "documents":
			i.eng.DeleteDocument(i.CurrentDB, i.CurrentColl, args[1])
>>>>>>> d4edcdab84f14de434df0b3539422c949ad05a23
		default:
			return fmt.Errorf("unkown argument for delete: %s", args[0])
		}*/
	return nil
}

func (i *Interpreter) cmdFind(rawQueries []string, args []string) error {
	fmt.Println("Comando find no implementado aún")
	return nil
}

func (i *Interpreter) cmdImport(args []string) error {
	fmt.Println("Comando import no implementado aún")
	return nil
}

func (i *Interpreter) cmdExport(args []string) error {
	fmt.Println("Comando export no implementado aún")
	return nil
}
