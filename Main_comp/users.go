package maincomp

import (
	"fmt"
)

// User interface defines the methods that any user type should implement
type User interface {
	GetUsername() string
	Authenticate(enteredPassword string) bool
	DisplayInfo()
}

// Base struct for User
type user struct {
	username string
	password string
}

// NewUser creates a new user instance
func NewUser(username, password string) *user {
	return &user{username: username, password: password}
}

// GetUsername returns the username of the user
func (u *user) GetUsername() string {
	return u.username
}

// Authenticate checks if the entered password is correct
func (u *user) Authenticate(enteredPassword string) bool {
	return u.password == enteredPassword
}

// AdminUser struct embeds the user struct
type AdminUser struct {
	*user
}

// NewAdminUser creates a new AdminUser instance
func NewAdminUser(username, password string) *AdminUser {
	return &AdminUser{NewUser(username, password)}
}

// DisplayInfo displays information for AdminUser
func (a *AdminUser) DisplayInfo() {
	fmt.Printf("Admin User: %s\n", a.username)
}

// StandardUser struct embeds the user struct
type StandardUser struct {
	*user
}

// NewStandardUser creates a new StandardUser instance
func NewStandardUser(username, password string) *StandardUser {
	return &StandardUser{NewUser(username, password)}
}

// DisplayInfo displays information for StandardUser
func (s *StandardUser) DisplayInfo() {
	fmt.Printf("Standard User: %s\n", s.username)
}
