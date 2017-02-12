// Package models - Category
package models

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/jsadwith/BestEver/database"
)

// Category - An category of items
type Category struct {
	gorm.Model         // Gets ID, CreatedAt, UpdatedAt, DeletedAt fields
	Name        string `json:"name"`
	Description string `json:"description"`
	User        User   `json:"user,omitempty"`
	Items       []Item `json:"items,omitempty"`
}

// AddCategory - Add Category
func AddCategory(category Category) (Category, error) {

	db := database.Client

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

// GetCategory - Retrieve a category
func GetCategory(id string, resources []string) (Category, error) {

	var category Category

	// Load a category
	db := database.Client
	query := db
	query = query.Where("ID = ?", id)

	// Expand related resources
	query = database.Expand(query, resources)

	if err := query.First(&category).Error; err != nil {
		errorStr := fmt.Sprintf("Error getting category (id=%s): %s", id, err)
		return Category{}, errors.New(errorStr)
	}

	return category, nil
}

// SearchCategories - Search categories by name
func SearchCategories(name string, resources []string) ([]Category, error) {

	var categories []Category

	db := database.Client
	query := db

	// Search for category by name (must use lower() for case insensitive)
	if name != "" && name != "all" {
		query = query.Where("LOWER(name) LIKE LOWER(?)", "%"+name+"%")
	}

	// Expand related resources
	query = database.Expand(query, resources)

	if err := query.Find(&categories).Error; err != nil {
		errorStr := fmt.Sprintf("Error searching categories: %s", err)
		return []Category{}, errors.New(errorStr)
	}

	return categories, nil
}
