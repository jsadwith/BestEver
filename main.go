/*

BestEver Prototype

*/

package main

import (
    "log"
    "net/http"

    "github.com/jsadwith/BestEver/config"
    "github.com/jsadwith/BestEver/database"
    "github.com/jsadwith/BestEver/models"
    "github.com/jsadwith/BestEver/routes"
)

func main() {

    // Load config
    config.Load()

    // Register routes, ensuring trailing slashes redirect to route - /route/ -> /route
    router := router.NewRouter()

    // Setup database connection and tables
    database.DbConnection = database.Connect()
    database.DbConnection.CreateTable(&models.Category{})
    database.DbConnection.CreateTable(&models.Item{})
    database.DbConnection.CreateTable(&models.User{})

    log.Printf("Serving on port " + config.Conf.App.Port)
    log.Fatal(http.ListenAndServe(":" + config.Conf.App.Port, router))
}
