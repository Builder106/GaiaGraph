package main

import (
    "fmt"
    "net/http"
    "os"
    "sort"

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

    // Extract and sort the country names
    var countryNames []string
    for name := range externalapis.Countries {
        countryNames = append(countryNames, name)
    }
    sort.Strings(countryNames)

    // Iterate over the sorted country names
    for _, name := range countryNames {
        lat, lon, err := externalapis.GetCountryGeocode(name, apiKey)
        if err != nil {
            fmt.Printf("Error fetching geocode for %s: %v\n", name, err)
            continue
        }
        fmt.Printf("%s: [latitude: %f, longitude: %f]\n", name, lat, lon)
    }

    http.HandleFunc("/climate-data", api.GetClimateData)

    fmt.Println("Server started at :8080")
    http.ListenAndServe(":8080", nil)
}