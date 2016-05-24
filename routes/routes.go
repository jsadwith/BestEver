/*

routes.go

Specifies routes

*/

package router

import (
    "github.com/jsadwith/BestEver/routes/api"
)

// Route type
type Route struct {
    Name        string
    Method      string
    Pattern     string
    JHandlerFunc JHandler
}

type Routes []Route

// API Routes
var routes = Routes {
    Route {
        "Add a new category",
        "POST",
        "/api/categories",
        api.AddCategory,
    },
    // TODO: update to {categoryId:[0-9]+}
    Route {
        "Get a category",
        "GET",
        "/api/categories/{categoryId}",
        api.GetCategory,
    },
    Route {
        "Search categories",
        "GET",
        "/api/categories/search/{query}",
        api.SearchCategories,
    },
    Route {
        "Add a new item to a category",
        "POST",
        "/api/items",
        api.AddItem,
    },
    Route {
        "Get an item",
        "GET",
        "/api/items/{itemId}",
        api.GetItem,
    },
    Route {
        "Search items in a category",
        "GET",
        "/api/items/search/{categoryId}/{query}",
        api.SearchItems,
    },

}
