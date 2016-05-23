/*

config.go

Sets up configuration

*/

package config

import (
    "os"
    "encoding/json"
)

type Config struct {
    App         App
    Database    Database
    Log         Log
}

type App struct {
    Hostname    string
    Port        string
}

type Database struct {
    Type        string
    Host        string
    Username    string
    Password    string
    Name        string
}

type Log struct {
    Loglocations  Loglocation
}

type Loglocation struct {
    Error       string
}

var (
    Conf *Config
)

func Load() {
    var file = "./config/config.json"
    _, err := os.Stat(file)
    if err != nil {
       panic("Config file cannot be found")
    }

    var config Config
    configFile, _ := os.Open(file)
    decoder := json.NewDecoder(configFile)

    if err := decoder.Decode(&config); err != nil {
        panic(err)
    }

    Conf = &config
}
