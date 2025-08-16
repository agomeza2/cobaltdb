package maincomp

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

// Relation struct represents a relationship between two nodes
type Relation struct {
	Source     *Node                  `json:"source"`     // Pointer to the source Node
	Target     *Node                  `json:"target"`     // Pointer to the target Node
	Name       string                 `json:"name"`       // Name of the relation
	Category   string                 `json:"category"`   // Category of the relation
	ID         int                    `json:"ID"`         // Unique ID for the relation
	Properties map[string]interface{} `json:"properties"` // Additional properties related to the relation
}

var globalIDCounterRelation int = 0 // Global ID counter for all nodes
var mu2 sync.Mutex

// Constructor for Relation
func NewRelation(source *Node, target *Node, category string, name string, properties ...interface{}) *Relation {
	mu2.Lock() // Lock to ensure thread safety
	defer mu2.Unlock()
	relation := &Relation{
		Source:     source,
		Target:     target,
		Category:   category,
		Name:       name,
		Properties: make(map[string]interface{}),
		ID:         globalIDCounterRelation,
	}
	globalIDCounterRelation++ // Get the next global ID for the relation
	relation.initializeProperties(properties...)
	return relation
}

// Initializes properties of the relation
func (r *Relation) initializeProperties(properties ...interface{}) {
	for i := 0; i < len(properties); i += 2 {
		if i+1 < len(properties) {
			key := properties[i].(string)
			value := properties[i+1]
			r.Properties[key] = value
		}
	}
}

// Prints the relation details
func (r *Relation) Show() {
	fmt.Printf("Relation (%s:%s) from %s to %s: ", r.Category, r.Name, r.Source.Name, r.Target.Name)
	for key, value := range r.Properties {
		fmt.Printf("%s: ", key)
		printElementRelation(value)
		fmt.Print(", ")
	}
	fmt.Println()
}

func (r *Relation) GetID() int {
	return r.ID
}
func (r *Relation) GetName() string {
	return r.Name
}
func (r *Relation) GetCategory() string {
	return r.Category
}
func (r *Relation) GetProperties() map[string]interface{} {
	return r.Properties
}
func (r *Relation) GetSource() *Node {
	return r.Source
}
func (r *Relation) GetTarget() *Node {
	return r.Target
}
func (r *Relation) Alter(key string, value interface{}) {
	r.Properties[key] = value
}

// Removes a property from the relation
func (r *Relation) Remove(key string) {
	delete(r.Properties, key)
}

// Adds a property to the relation
func (r *Relation) Add(key string, value interface{}) {
	r.Properties[key] = value
}

// Converts the relation to JSON
func (r *Relation) ToJSON() ([]byte, error) {
	data := map[string]interface{}{
		"RelationID": r.ID,
		"category":   r.Category,
		"name":       r.Name,
		"properties": r.Properties,
		"source":     r.Source.ID,
		"target":     r.Target.ID,
	}

	return json.MarshalIndent(data, "", "  ")
}

// Writes the relation attributes to a JSON file
func (r *Relation) WriteToJsonFile(dbPath, filename string) {
	// Create the Relations directory if it doesn't exist
	relationsDir := filepath.Join(dbPath, "Relations")
	if err := os.MkdirAll(relationsDir, os.ModePerm); err != nil {
		fmt.Printf("Failed to create directory: %v\n", err)
		return
	}

	filePath := filepath.Join(relationsDir, fmt.Sprintf("%s-%d.json", filename, r.ID))
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("Failed to open file: %v\n", err)
		return
	}
	defer file.Close()

	jsonData, err := r.ToJSON() // Serialize the relation to JSON
	if err != nil {
		fmt.Printf("Failed to convert relation to JSON: %v\n", err)
		return
	}

	// Write the JSON data to the file
	if _, err := file.Write(jsonData); err != nil {
		fmt.Printf("Failed to write to file: %v\n", err)
		return
	}
	fmt.Printf("Relation attributes written to %s\n", filePath)
}

// Helper function to print elements of different types
func printElementRelation(element interface{}) {
	switch v := element.(type) {
	case string:
		fmt.Printf("\"%s\"", v)
	case int:
		fmt.Print(v)
	case float64:
		fmt.Print(v)
	case bool:
		fmt.Print(v)
	default:
		fmt.Println("Unsupported type")
	}
}
