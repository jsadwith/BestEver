/*

item.go

Items

*/

package models

import (
    "github.com/jsadwith/BestEver/database"
    "github.com/jinzhu/gorm"
)

// Item
type Item struct {
    gorm.Model    // Gets ID, CreatedAt, UpdatedAt, DeletedAt fields
    Name          string     `json:"name"`
    Description   string     `json:"description"`
    User          User       `json:"user,omitempty"`
    CategoryID    uint       `json:"category"`
}

// Add Item
func AddItem(item Item) Item {
    // Add Item to DB
    db := database.DbConnection
    db.Create(&item)

    return item
}

// Retrieve item
// TODO: Add filtering
// TODO: MAKE THIS WORK
func GetItem(id int32) int32 {
    // Load all categories with items
    db := database.DbConnection
    db.Find(&id)

    return id
}
