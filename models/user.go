/*

author.go

Creator of stuff

*/

package models

import (
    "github.com/jsadwith/BestEver/database"
    "github.com/jinzhu/gorm"
)

// User
type User struct {
    gorm.Model    // Gets ID, CreatedAt, UpdatedAt, DeletedAt fields
    Name          Name       `json:"name"`
    Email         string     `json:"email"`
}

// Name
type Name struct {
    First         string    `json:"first"`
    Last          string    `json:"last"`
    Display       string    `json:"display"`
}

// Add User
func AddUser(user User) User {
    // Add User to DB
    db := database.DbConnection
    db.Create(&user)

    return user
}

// Retrieve user
// TODO: Add filtering
// TODO: MAKE THIS WORK
func GetUser(id int32) int32 {
    db := database.DbConnection
    db.Find(&id)

    return id
}
