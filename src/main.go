package main

import (
	"fmt"
	"path/filepath"
	"github.com/agomeza2/cobalt-g/Main_comp"
	"github.com/agomeza2/cobalt-g/Storage"
)

func main() {
	storage := &storage.Storage{}

	// Example usage
	storage.CreateUser("Alice")
	storage.CreateDB("test", "Alice")

	dbPath:="../db/Alice/test"
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
