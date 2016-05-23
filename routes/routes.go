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
        "/api/category",
        api.AddCategory,
    },
    // Route {
    //     "Add a new item to a category",
    //     "POST",
    //     "/api/item/add",
    //     api.AddItem,
    // },
    Route {
        "Get a category",
        "GET",
        "/api/category/{categoryId}",
        api.GetCategory,
    },
    // Route {
    //     "Get a category with items",
    //     "GET",
    //     "/api/category/{categoryId}/full",
    //     api.GetCategoryFull,
    // },
}
