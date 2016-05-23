/*

database.go

Handles setup of database

*/

package database

import (
    "github.com/jsadwith/BestEver/config"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
    DbConnection *gorm.DB
)

// Connect to the database
// NOTE: Only supports Postgres at this time
func Connect() *gorm.DB {

    db, err := gorm.Open("postgres",
        "host=" + config.Conf.Database.Host +
        " user=" + config.Conf.Database.Username +
        " dbname=" + config.Conf.Database.Name +
        " sslmode=disable" +
        " password=" + config.Conf.Database.Password)

    db.LogMode(config.Conf.Database.Log)

    if err != nil {
        panic("Failed to connect to db")
    }

    return db
}
