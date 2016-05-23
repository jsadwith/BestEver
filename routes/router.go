/*

router.go

Implements routes from routes.go

*/
package router

import (
    "fmt"
    "net/http"

    "github.com/jsadwith/BestEver/util"
    "github.com/gorilla/mux"
)

func NewRouter() *mux.Router {

    // Create a new Gorilla Mux router
    router := mux.NewRouter().StrictSlash(true)

    // ...with routes defined below
    for _, route := range routes {
        var handler http.Handler

        handler = route.JHandlerFunc
        handler = util.RouteLogger(handler, route.Name)

        router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(handler)
    }

    return router
}

// Custom HTTP Handler so we can return error codes and errors
type JHandler func(http.ResponseWriter, *http.Request) (int, error)

func (fn JHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if status, err := fn(w, r); err != nil {
        // Log error
        logContent := fmt.Sprintf("HTTP Error (%s): %+v", r.RequestURI, err)
        util.FileLog("error", logContent)

        // Return error based on status code
        switch status {
            // TODO: Add additional status codes
            // case http.StatusInternalServerError:
            default:
                // Catch any other errors we haven't explicitly handled
                http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
        }
    } else {
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    }
}
