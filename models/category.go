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
    User          User       `json:"user,omitempty"`
    Items         []Item     `json:"items,omitempty"`
}

// Add Category
func AddCategory(category Category) (Category, error) {
    // Add Item to DB
    db := database.DbConnection

    // Check if category exists
    var count int
    db.Where("name = ?", category.Name).Find(&category).Count(&count)
    if count > 0 {
        errorStr := fmt.Sprintf("Category already exists (%s)", category.Name)
        return Category{}, errors.New(errorStr)
    }

    // Create and check for errors
    if err := db.Create(&category).Error; err != nil {
        errorStr := fmt.Sprintf("Error creating category (%s): %s", category.Name, err)
        return Category{}, errors.New(errorStr)
    }

    return category, nil
}

// Retrieve a category
func GetCategory(id uint) (Category, error) {
    var category Category

    // Load a category
    db := database.DbConnection
    if err := db.Where("ID = ?", id).First(&category).Error; err != nil {
        errorStr := fmt.Sprintf("Error getting category (id=%s): %s", id, err)
        return Category{}, errors.New(errorStr)
    }

    return category, nil
}

// Search categories by name
func SearchCategories(name string) ([]Category, error) {
    var categories []Category
    var err error

    // Load all categories with items
    db := database.DbConnection

    // Search for category by name (must use lower() for case insensitive)
    if name != "" && name != "all" {
        err = db.Where("LOWER(name) LIKE LOWER(?)", "%" + name + "%").Find(&categories).Error
    } else { // If no search supplied, just grab all
        err = db.Preload("Items").Find(&categories).Error
    }

    if err != nil {
        errorStr := fmt.Sprintf("Error searching categories: %s", err)
        return []Category{}, errors.New(errorStr)
    }

    return categories, nil
}