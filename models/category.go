/*

category.go

Categories

*/

package models

import (
    "errors"
    "fmt"

    "github.com/jsadwith/BestEver/database"
    "github.com/jinzhu/gorm"
)

// Category
type Category struct {
    gorm.Model    // Gets ID, CreatedAt, UpdatedAt, DeletedAt fields
    Name          string     `json:"name"`
    Description   string     `json:"description"`
    User          User       `json:"user"`
    Items         []Item     `json:"items"`
}

// Add Category
func AddCategory(category Category) (Category, error) {
    // Add Item to DB
    db := database.DbConnection

    // Check if category exists
    if db.NewRecord(category) == false {
        errorStr := fmt.Sprintf("Category already exists (%s)", category.Name)
        return Category{}, errors.New(errorStr)
    }

    if err := db.Create(&category).Error; err != nil {
        errorStr := fmt.Sprintf("Error creating category (%s): %s", category.Name, err)
        return Category{}, errors.New(errorStr)
    }

    return category, nil
}

// Retrieve categories
// TODO: Add filtering
func GetCategories() []Category {
    var categories []Category

    // Load all categories with items
    db := database.DbConnection
    db.Preload("Items").Find(&categories)

    return categories
}
