package main

import (
    "log"
    "net/http"
    "GaiaGraph/internal/api"
    "GaiaGraph/internal/db"
)

func main() {
    db.InitDB("user=youruser dbname=yourdb sslmode=disable")

    http.HandleFunc("/climate-data", api.GetClimateData)

    log.Println("Server started at :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}