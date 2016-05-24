/*

category.go

Handles storage/update of categories

*/

package api

import (
    "encoding/json"
    "errors"
    "io"
    "io/ioutil"
    "log"
    "net/http"
    "strings"

    "github.com/jsadwith/BestEver/models"
    "github.com/gorilla/mux"
)

// Add Category
func AddCategory(w http.ResponseWriter, r *http.Request) (int, error) {

    // Open request body
    // Set limit to prevent large JSON POSTs
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
        return 500, err
    }

    if err := r.Body.Close(); err != nil {
        return 500, err
    }

    var category models.Category

    if err := json.Unmarshal(body, &category); err != nil {
        return 500, err
    }

    // TODO: Add more validation
    if category.Name == "" {
        return 500, errors.New("Category name empty")
    }

    // Insert category
    category, err = models.AddCategory(category)
    if err != nil {
        return 500, err
    }

    log.Printf("api.Category - Add Category: Category created (%s)", category.Name)

    // Return entity as JSON
    if err := json.NewEncoder(w).Encode(category); err != nil {
        return 500, err
    }

    return 200, nil
}

// Get category
func GetCategory(w http.ResponseWriter, r *http.Request) (int, error) {

    vars := mux.Vars(r)

    // Get related resources
    resourcesGet := r.URL.Query().Get("expand")
    resources := strings.Split(resourcesGet, ",")

    // Get category
    category, err := models.GetCategory(vars["categoryId"], resources)
    if err != nil {
        return 500, err
    }

    // Return entity as JSON
    if err := json.NewEncoder(w).Encode(category); err != nil {
        return 500, err
    }

    return 200, nil
}

// Search categories
func SearchCategories(w http.ResponseWriter, r *http.Request) (int, error) {

    vars := mux.Vars(r)

    // Get related resources
    resourcesGet := r.URL.Query().Get("expand")
    resources := strings.Split(resourcesGet, ",")

    // Search
    categories, err := models.SearchCategories(vars["query"], resources)
    if err != nil {
        return 500, err
    }

    // Return entity as JSON
    if err := json.NewEncoder(w).Encode(categories); err != nil {
        return 500, err
    }

    return 200, nil
}

func UpdateCategory(w http.ResponseWriter, r *http.Request) (int, error) {
    return 200, nil
}

