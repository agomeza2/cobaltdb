package storage

import (
	"fmt"
	"os"
	"path/filepath"
)

// Storage struct represents the storage system
type Storage struct{}

// createFolder creates a directory at the specified path
func (s *Storage) createFolder(dbName string, path string) {
	if dbName != "" {
		path = filepath.Join(path, dbName)
	}

	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		fmt.Printf("Failed to create folder: %s\n", path)
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Folder created successfully: %s\n", path)
}

// createDB creates a database structure with Nodes and Relations folders
func (s *Storage) CreateDB(dbName, path string, user string) {
	userPath := filepath.Join(user, dbName)
	s.createFolder(dbName, userPath)

	nodesPath := filepath.Join(userPath, "Nodes")
	relationsPath := filepath.Join(userPath, "Relations")
	s.createFolder(nodesPath, userPath)
	s.createFolder(relationsPath, userPath)
}

// createUser creates a user directory
func (s *Storage) CreateUser(userName string, path string) {
	s.createFolder(userName, path)
}
