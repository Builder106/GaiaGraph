package main

import (
    "fmt"
    "net/http"
    "os"

    "GaiaGraph/internal/api"
    "GaiaGraph/internal/db"
    "GaiaGraph/externalapis"
)

func main() {
    dataSourceName := "user=olayinkav password=T0p0l0gic^ls dbname=climate_data sslmode=disable"
    db.InitDB(dataSourceName)

    externalapis.LoadEnv()
    apiKey := os.Getenv("GEOAPIFY_API_KEY")
    if apiKey == "" {
        fmt.Println("API_KEY not set in .env file")
        return
    }

    http.HandleFunc("/climate-data", api.GetCombinedData)

    fmt.Println("Server started at :8080")
    http.ListenAndServe(":8080", nil)
}