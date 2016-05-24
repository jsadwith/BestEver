/*

item.go

Items

*/

package models

import (
    "errors"
    "fmt"

    "github.com/jsadwith/BestEver/database"
    "github.com/jinzhu/gorm"
)

// Item
type Item struct {
    gorm.Model    // Gets ID, CreatedAt, UpdatedAt, DeletedAt fields
    Name          string     `json:"name"`
    Description   string     `json:"description"`
    User          User       `json:"user,omitempty"`
    CategoryID    uint       `json:"category_id"`
}

// Add Item
func AddItem(item Item) (Item, error) {

    db := database.DbConnection

    // Check if item exists in this category
    var count int
    db.Where("name = ? AND category_id = ?", item.Name, item.CategoryID).Find(&item).Count(&count)
    if count > 0 {
        errorStr := fmt.Sprintf("Item already exists (%s)", item.Name)
        return Item{}, errors.New(errorStr)
    }

    // Create and check for errors
    if err := db.Create(&item).Error; err != nil {
        errorStr := fmt.Sprintf("Error creating item (%s): %s", item.Name, err)
        return Item{}, errors.New(errorStr)
    }

    return item, nil
}

// Retrieve an item
func GetItem(id string) (Item, error) {
    var item Item

    // Load an item
    db := database.DbConnection
    if err := db.Where("ID = ?", id).First(&item).Error; err != nil {
        errorStr := fmt.Sprintf("Error getting item (id=%s): %s", id, err)
        return Item{}, errors.New(errorStr)
    }

    return item, nil
}

// Search items by name
func SearchItems(categoryId string, name string) ([]Item, error) {
    var items []Item
    var err error

    db := database.DbConnection

    // Require categoryId
    if categoryId == "" {
        return []Item{}, errors.New("Category ID required")
    }

    // Search for item by name (must use lower() for case insensitive)
    if name != "" && name != "all" {
        err = db.Where("category_id = ? AND LOWER(name) LIKE LOWER(?)", categoryId, "%" + name + "%").Find(&items).Error
    } else { // If no search supplied, grab all in this category
        err = db.Where("category_id = ?", categoryId).Find(&items).Error
    }

    if err != nil {
        errorStr := fmt.Sprintf("Error searching categories: %s", err)
        return []Item{}, errors.New(errorStr)
    }

    return items, nil
}