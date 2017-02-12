// Package api - Handles storage/update of items
package api

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jsadwith/BestEver/models"
)

// AddItem - Add Item
func AddItem(w http.ResponseWriter, r *http.Request) (int, error) {

	// Open request body
	// Set limit to prevent large JSON POSTs
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		return 500, err
	}

	if err := r.Body.Close(); err != nil {
		return 500, err
	}

	var item models.Item

	if err := json.Unmarshal(body, &item); err != nil {
		return 500, err
	}

	// TODO: Add more validation
	// Verify item name is set
	if item.Name == "" {
		return 500, errors.New("Item name empty")
	} else if item.CategoryID < 1 { // Verify category ID is set
		return 500, errors.New("Category ID required")
	}

	// Insert item
	item, err = models.AddItem(item)
	if err != nil {
		return 500, err
	}

	log.Printf("api.Item - Add Item: Item created (%s) in category %b", item.Name, item.CategoryID)

	// Return entity as JSON
	if err := json.NewEncoder(w).Encode(item); err != nil {
		return 500, err
	}

	return 200, nil
}

// GetItem - Get item
func GetItem(w http.ResponseWriter, r *http.Request) (int, error) {

	vars := mux.Vars(r)

	// Get item
	item, err := models.GetItem(vars["itemId"])
	if err != nil {
		return 500, err
	}

	// Return entity as JSON
	if err := json.NewEncoder(w).Encode(item); err != nil {
		return 500, err
	}

	return 200, nil
}

// SearchItems - Search items
func SearchItems(w http.ResponseWriter, r *http.Request) (int, error) {

	vars := mux.Vars(r)

	// Search
	items, err := models.SearchItems(vars["categoryId"], vars["query"])
	if err != nil {
		return 500, err
	}

	// Return entity as JSON
	if err := json.NewEncoder(w).Encode(items); err != nil {
		return 500, err
	}

	return 200, nil
}

// UpdateItem - Update items
func UpdateItem(w http.ResponseWriter, r *http.Request) (int, error) {
	return 200, nil
}
