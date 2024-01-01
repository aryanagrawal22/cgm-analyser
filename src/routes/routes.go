// routes/routes.go
package routes

import (
    "net/http"
    "github.com/aryanagrawal22/cgm-analyser/controllers"
)

func InitRoutes() {
    http.HandleFunc("/", controllers.HandleRoot)
}