package main

import (
    "log"
    "net/http"
    "GaiaGraph/internal/api"
    "GaiaGraph/internal/db"
)

func main() {
    dataSourceName := "user=olayinkav password=T0p0l0gic^ls dbname=climate_data sslmode=disable"
    db.InitDB(dataSourceName)

    http.HandleFunc("/climate-data", api.GetClimateData)

    log.Println("Server started at :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}