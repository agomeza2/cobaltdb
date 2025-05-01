package maincomp

import (
    "encoding/json"
    "fmt"
    "os"
    "path/filepath"
    "sync"
)

// Node represents a node with unique properties
type Node struct {
    Name       string                 `json:"name"`
    Category   string                 `json:"category"`
    ID         int                    `json:"ID"`
    Properties map[string]interface{} `json:"properties"`
}

// Global counters
var categoryCounters = make(map[string]int)
var globalIDCounter int // Global ID counter for all nodes
var mu sync.Mutex      // Mutex to synchronize access

// Constructor for Node
func NewNode(category, name string, properties ...interface{}) *Node {
    mu.Lock() // Lock to ensure thread safety
    defer mu.Unlock()

    node := &Node{
        Category:   category,
        Name:       name,
        Properties: make(map[string]interface{}),
        ID:         globalIDCounter, // Assign global ID
    }
    globalIDCounter++ // Increment global ID counter

    node.initializeProperties(properties...)
    return node
}

// Initializes properties of the node
func (n *Node) initializeProperties(properties ...interface{}) {
    for i := 0; i < len(properties); i += 2 {
        if i+1 < len(properties) {
            n.Properties[properties[i].(string)] = properties[i+1]
        }
    }
}

// Prints the node properties
func (n *Node) Show() {
    fmt.Printf("(%s:%s)(", n.Category, n.Name)
    count := 0
    totalProps := len(n.Properties)

    for key, value := range n.Properties {
        fmt.Printf("%s:", key)
        printElementNode(value)

        count++
        if count < totalProps {
            fmt.Print(",")
        }
    }
    fmt.Println(")")
}

// Alters a property
func (n *Node) Alter(key string, value interface{}) {
    n.Properties[key] = value
}

// Alters the name of the node
func (n *Node) AlterName(name string) {
    n.Name = name
}

// Removes a property
func (n *Node) Remove(key string) {
    delete(n.Properties, key)
}

// Adds a property
func (n *Node) Add(key string, value interface{}) {
    n.Properties[key] = value
}

// Converts the node to JSON
func (n *Node) ToJSON() ([]byte, error) {
    data :=map[string]interface{}{
        "ID":    n.ID, 
        "category":   n.Category,                                                                   
        "name":       n.Name,
        "properties": n.Properties,
    }

    return json.MarshalIndent(data, "", "  ")
}

// Writes the node attributes to a JSON file
func (n *Node) WriteToJsonFile(dbPath, filename string) {
    // Create the Nodes directory if it doesn't exist
    nodesDir := filepath.Join(dbPath, "Nodes")
    err := os.MkdirAll(nodesDir, os.ModePerm)
    if err != nil {
        fmt.Printf("Failed to create directory: %v\n", err)
        return
    }

    ID := fmt.Sprintf("%d", n.ID)
    filePath := filepath.Join(nodesDir, fmt.Sprintf("%s-%s.json", filename, ID))
    file, err := os.Create(filePath)
    if err != nil {
        fmt.Printf("Failed to open file: %v\n", err)
        return
    }
    defer file.Close()

    jsonData, err := n.ToJSON() // Serialize the node to JSON
    if err != nil {
        fmt.Printf("Failed to convert node to JSON: %v\n", err)
        return
    }

    // Write the JSON data to the file
    _, err = file.Write(jsonData)
    if err != nil {
        fmt.Printf("Failed to write to file: %v\n", err)
        return
    }
    fmt.Printf("Node attributes written to %s\n", filePath)
}

// Helper function to print elements of different types
func printElementNode(element interface{}) {
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
