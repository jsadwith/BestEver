// package Main - BestEver prototype

package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/jsadwith/BestEver/database"
	"github.com/jsadwith/BestEver/models"
	"github.com/jsadwith/BestEver/routes"
)

var (
	port = flag.String(
		"PORT",
		os.Getenv("PORT"),
		"",
	)
)

func main() {
	// Register routes, ensuring trailing slashes redirect to route - /route/ -> /route
	router := router.NewRouter()

	err := database.Connect()

	if err != nil {
		log.Fatalf("Error connecting to Postgres: %s", err)
	}

	// Table creation and auto migrate
	database.Client.CreateTable(&models.Category{})
	database.Client.CreateTable(&models.Item{})
	database.Client.CreateTable(&models.User{})
	database.Client.AutoMigrate(&models.Category{}, &models.Item{}, &models.User{})

	log.Printf("Serving on port " + *port)
	log.Fatal(http.ListenAndServe(":"+*port, router))
}
