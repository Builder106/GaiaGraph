package main

import (
    "fmt"
    "log"
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
        log.Fatal("API_KEY not set in .env file")
    }

    countryGeocodes := make(map[string]map[string]float64)

    for name, code := range externalapis.Countries {
        lat, lon, err := externalapis.GetCountryGeocode(name, apiKey)
        if err != nil {
            log.Printf("Error fetching geocode for %s: %v", name, err)
            continue
        }
        countryGeocodes[code] = map[string]float64{
            "latitude":  lat,
            "longitude": lon,
        }
    }

    for code, geocode := range countryGeocodes {
        fmt.Printf("%s: %v\n", code, geocode)
    }

    http.HandleFunc("/climate-data", api.GetClimateData)

    log.Println("Server started at :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}