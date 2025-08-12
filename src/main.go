//package main

/*
import (

	"fmt"
	"path/filepath"

	maincomp "github.com/agomeza2/cobalt-g/Main_comp"
	storage "github.com/agomeza2/cobalt-g/Storage"

)

	func main() {
		storage := &storage.Storage{}

		// Example usage
		storage.CreateUser("Alice")
		storage.CreateDB("test", "Alice")

		dbPath := "../db/Alice/test"

			// Create some nodes
			node1 := maincomp.NewNode("Person", "Alice", "age", 30, "city", "Wonderland")
			node2 := maincomp.NewNode("Person", "Bob", "age", 25, "city", "Builderland")
			node3 := maincomp.NewNode("Animal", "Charlie", "type", "Dog")

			// Display nodes
			node1.Show()
			node2.Show()
			node3.Show()

			// Create relations between nodes
			relation1 := maincomp.NewRelation(node1, node2, "Friendship", "Close Friends", "duration", 5)
			relation2 := maincomp.NewRelation(node2, node3, "Ownership", "Owner", "since", 2020)
			relation3 := maincomp.NewRelation(node3, node2, "Friendship", "Close Friends", "duration", 5)
			relation4 := maincomp.NewRelation(node1, node3, "Ownership", "Owner", "since", 2020)
			// Display relations
			relation1.Show()
			relation2.Show()

			// Write nodes to JSON files
			node1.WriteToJsonFile(dbPath, node1.Name) // Adjust dbPath as needed
			node2.WriteToJsonFile(dbPath, node2.Name)
			node3.WriteToJsonFile(dbPath, node3.Name )

			// Write relations to JSON files
			relation1.WriteToJsonFile(dbPath, relation1.Name) // Write relation1 to file
			relation2.WriteToJsonFile(dbPath, relation2.Name) // Write relation2 to file
			relation3.WriteToJsonFile(dbPath, relation3.Name) // Write relation1 to file
			relation4.WriteToJsonFile(dbPath, relation4.Name) // Write relation2 to file


		// Create a map to hold users
		users := make(map[string]maincomp.User)

		// Add users to the map
		users["admin"] = maincomp.NewAdminUser("admin", "admin123")
		users["user1"] = maincomp.NewStandardUser("user1", "user123")
		users["user2"] = maincomp.NewStandardUser("user2", "pass123")

		var username, password string

		// Prompt for username and password
		fmt.Print("Enter username: ")
		fmt.Scan(&username)
		fmt.Print("Enter password: ")
		fmt.Scan(&password)

		// Authenticate the user
		if user, exists := users[username]; exists && user.Authenticate(password) {
			fmt.Println("Authentication successful.")
			user.DisplayInfo()
		} else {
			fmt.Println("Incorrect username or password.")
		}
		dp := &maincomp.DataProcessor{} // Correctly reference the package

		// Define the path to the Excel file
		filePath := filepath.Join("../import", "salary_country.xlsx")

		// Process the nodes from the Excel file
		fmt.Println("Starting to process nodes...")
		dp.ProcessDataToNodeExcel(filePath)

		// Show the nodes
		fmt.Println("Showing nodes...")
		dp.ShowNodes()
		dp.SaveNodes(dbPath)

		// Process relations between nodes
		fmt.Println("Processing relations...")
		dp.ProcessDataToRelationExcel()

		fmt.Println("Processing completed.")
		dp.ShowRelations()
		dp.SaveRelations(dbPath)
	}

	func main() {
		storage := &storage.Storage{}

		// Example usage
		storage.CreateUser("Alice")
		storage.CreateDB("test", "Alice")

		parser, err := l27.NewParser()
		if err != nil {
			panic(err)
		}
		i:= l27.NewInterpreter(storage, parser)
		dp := &maincomp.DataProcessor{}
		filePath := filepath.Join("../import", "salary_country.xlsx")

		// Process the nodes from the Excel file
		fmt.Println("Starting to process nodes...")
		dp.ProcessDataToNodeExcel(filePath)



		// Mensaje de inicio
		fmt.Println("CobaltDB CLI (versión demo) — escribe 'exit' o 'quit' para salir.")
		fmt.Println("Ejemplo: CREATE DATABASE test; SELECT DATABASE test; CREATE NODE {category:people,name:\"Jo\",age:20}")

		// REPL simple que acepta múltiples sentencias separadas por ';'
		reader := bufio.NewReader(os.Stdin)
		for {
			fmt.Print("> ")
			line, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error lectura:", err)
				return
			}
			line = strings.TrimSpace(line)
			if line == "" {
				continue
			}
			lower := strings.ToLower(line)
			if lower == "exit" || lower == "quit" {
				fmt.Println("Adiós.")
				return
			}
			// permite múltiples sentencias separadas por ';'
			parts := splitStatements(line)
			for _, stmtText := range parts {
				if strings.TrimSpace(stmtText) == "" {
					continue
				}
				prog, err := parser.ParseString("", stmtText)
				if err != nil {
					fmt.Println("Error de parseo:", err)
					continue
				}
				for _, st := range prog.Stmts {
					i.Execute(st)
				}
			}
		}
	}

	func splitStatements(s string) []string {
		// split por ';' respetando que no estamos manejando ; dentro de strings (básico)
		raw := strings.Split(s, ";")
		out := []string{}
		for _, r := range raw {
			if t := strings.TrimSpace(r); t != "" {
				out = append(out, t)
			}
		}
		return out
	}
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/alecthomas/participle/v2"

	l27 "cobaltdb-local/L-27" // Ajusta al import real de tu proyecto
)

func main() {
	// Instancia única de GraphDB para toda la sesión
	db := l27.NewGraphDB()

	// Construir parser con tu lexer
	parser, err := participle.Build[l27.Program](
		participle.Lexer(l27.LexerDef),        // Definido en lexer.go
		participle.Unquote("String"),          // Manejo de strings sin comillas
		participle.CaseInsensitive("Keyword"), // Palabras clave sin distinción de mayúsculas
	)
	if err != nil {
		panic(fmt.Errorf("error construyendo parser: %w", err))
	}

	fmt.Println("CobaltDB CLI — escribe 'exit' o 'quit' para salir")
	fmt.Println("Ejemplo: CREATE DATABASE test; SELECT DATABASE test; CREATE NODE {category:people,name:\"Jo\",age:20}")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(">> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error de lectura:", err)
			return
		}

		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		// Comandos para salir
		if strings.EqualFold(line, "exit") || strings.EqualFold(line, "quit") {
			fmt.Println("Adiós.")
			return
		}

		// Permitir múltiples sentencias separadas por ";"
		statements := splitStatements(line)
		for _, stmtText := range statements {
			if stmtText == "" {
				continue
			}

			ast, err := parser.ParseString("", stmtText)
			if err != nil {
				fmt.Println("Error de parseo:", err)
				continue
			}

			// Ejecutar cada sentencia individual
			for _, st := range ast.Stmts {
				db.Execute(st) // Ejecuta directamente, o cambia si quieres devolver error
			}
		}
	}
}

// splitStatements divide sentencias separadas por ";" ignorando espacios
func splitStatements(input string) []string {
	raw := strings.Split(input, ";")
	out := make([]string, 0, len(raw))
	for _, stmt := range raw {
		if t := strings.TrimSpace(stmt); t != "" {
			out = append(out, t)
		}
	}
	return out
}
