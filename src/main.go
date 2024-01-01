package main

import (
    "log"
    "net/http"
    "github.com/aryanagrawal22/cgm-analyser/routes"
    "github.com/aryanagrawal22/cgm-analyser/crons"
    "github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    routes.InitRoutes()

    go crons.StartCron()

    port := ":9876"
    log.Printf("Server started on port %s\n", port)
    if err := http.ListenAndServe(port, nil); err != nil {
        log.Fatal(err)
    }
}