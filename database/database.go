/*

database.go

Handles setup of database

*/

package database

import (
	"flag"
	"log"
	"os"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Postgres dialect for gorm
)

// Client - Global DB connection
var (
	Client *gorm.DB
)

var (
	host = flag.String(
		"DB_POSTGRES_HOST",
		os.Getenv("DB_POSTGRES_HOST"),
		"",
	)
	user = flag.String(
		"DB_POSTGRES_USER",
		os.Getenv("DB_POSTGRES_USER"),
		"",
	)
	password = flag.String(
		"DB_POSTGRES_PASSWORD",
		os.Getenv("DB_POSTGRES_PASSWORD"),
		"",
	)
	dbName = flag.String(
		"DB_POSTGRES_NAME",
		os.Getenv("DB_POSTGRES_NAME"),
		"",
	)
	logMode = flag.String(
		"DB_POSTGRES_LOG",
		os.Getenv("DB_POSTGRES_LOG"),
		"",
	)
)

// Connect - Connect to the database
func Connect() error {

	connectionString := "host=" + *host +
		" user=" + *user +
		" dbname=" + *dbName +
		" sslmode=disable" +
		" password=" + *password

	log.Printf("Connecting to %s", connectionString)
	connection, err := gorm.Open("postgres", connectionString)

	log.Printf("past connection")
	logModeBool, _ := strconv.ParseBool(*logMode)
	connection.LogMode(logModeBool)

	if err != nil {
		panic("Failed to connect to db")
	}

	Client = connection

	return nil
}

// Expand - Expand the query to additional resources
func Expand(query *gorm.DB, resources []string) *gorm.DB {
	for i := range resources {
		if resources[i] != "" {
			query = query.Preload(resources[i])
		}
	}
	return query
}
