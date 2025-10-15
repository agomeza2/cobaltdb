package storage

import (
	"fmt"
	"log"
	"strconv"
	"github.com/xuri/excelize/v2"
	comp "cobaltdb-local/Main_comp"
)

type DataProcessor struct {
	Nodes      [] *comp.Node	
	Relations  [] *comp.Relation 
}


// ProcessDataToNodeExcel reads nodes from an Excel file
func (dp *DataProcessor) ProcessDataToNodeExcel(filePath string) {
	fmt.Println("Processing Excel file:", filePath)

	f, err := excelize.OpenFile(filePath)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}

	// Get all rows in the first sheet
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		log.Fatalf("Error getting rows: %v", err)
	}

	var headers []string
	firstRow := true

	for _, row := range rows {
		if firstRow {
			headers = row
			firstRow = false
			continue
		}

		if len(row) < len(headers) {
			log.Printf("Row does not have enough columns: %d\n", len(row))
			continue
		}

		name := row[0]
		category := headers[0]
		node := comp.NewNode(category, name) // Assuming NewNode returns a pointer

		for i := 1; i < len(row); i++ {
			node.Add(headers[i], row[i])
		}
		dp.Nodes = append(dp.Nodes,node) 
		comp.AddNode(node)  // Append node (which is a pointer)
	}
}

// ProcessDataToRelationExcel processes relations based on node attributes
func (dp *DataProcessor) ProcessDataToRelationExcel() {
	count := 0
	// Iterate through each node
	for i := 0; i < len(dp.Nodes); i++ {
		for j := i + 1; j < len(dp.Nodes); j++ {
			node1 := dp.Nodes[i]
			node2 := dp.Nodes[j]

			// Search for common attributes between the two nodes
			for key, val1 := range node1.Properties {
				if val2, exists := node2.Properties[key]; exists && val1 == val2 {
					// Create the relation using the NewRelation constructor
					relation := comp.NewRelation(
						node1,            // Source Node
						node2,            // Destination Node
						"RelatedBy_"+key, // Category
						"CommonAttribute"+key+"_"+strconv.Itoa(count), // Name
						key, val1, // Properties (key-value pair)
					)

					// Add the relation to the list of relations
					dp.Relations=append(dp.Relations,relation) 
					comp.AddRelation(relation)

					// Increment count for each relation created
					count++
				}
			}
		}
	}
}
