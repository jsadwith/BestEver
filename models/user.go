// Package models - User
package models

import (
	"github.com/jinzhu/gorm"
	"github.com/jsadwith/BestEver/database"
)

// User - End user
type User struct {
	gorm.Model        // Gets ID, CreatedAt, UpdatedAt, DeletedAt fields
	Name       Name   `json:"name"`
	Email      string `json:"email"`
}

// Name - End user name
type Name struct {
	First   string `json:"first"`
	Last    string `json:"last"`
	Display string `json:"display"`
}

// AddUser - Add User
func AddUser(user User) User {
	// Add User to DB
	db := database.Client
	db.Create(&user)

	return user
}

// GetUser - Retrieve user
// TODO: Add filtering
// TODO: MAKE THIS WORK
func GetUser(id int32) int32 {
	db := database.Client
	db.Find(&id)

	return id
}
