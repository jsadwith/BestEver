/*

logger.go

Logging!

*/

package util

import (
    "io/ioutil"
    "log"
    "net/http"
    "strings"
    "time"

    "github.com/jsadwith/BestEver/config"
)

// Logger to log all HTTP requests
func RouteLogger(inner http.Handler, name string) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()

        inner.ServeHTTP(w, r)

        log.Printf(
            "%s\t%s\t%s\t%s",
            r.Method,
            r.RequestURI,
            name,
            time.Since(start),
        )
    })
}

// Log to file
func FileLog(logType string, logContent string) error {

    var fileLocation string

    switch logType {
        // TODO: Prepend errors with some sort of timestamp
        case "error":
            // Also send to app log
            log.Printf(logContent)
            fileLocation = config.Conf.Log.Loglocations.Error
        default:
            // Also send to app log
            log.Printf(logContent)
            fileLocation = config.Conf.Log.Loglocations.Error
    }

    t := time.Now()

    // Replace macros and write to file
    fileLocation = strings.Replace(fileLocation, "@date@", t.Format("20060102"), 1)
    err := ioutil.WriteFile(fileLocation, []byte(logContent + "\n"), 0644)
    if err != nil {
        return err
    }

    return nil
}