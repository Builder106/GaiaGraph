package externalapis

import (
    "encoding/json"
    "fmt"
    "io"
    "log"
    "net/http"

    "github.com/joho/godotenv"
)

type Geocode struct {
    Lat float64 `json:"lat"`
    Lon float64 `json:"lon"`
}

func LoadEnv() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
}

func GetCountryGeocode(countryName, apiKey string) (float64, float64, error) {
    url := fmt.Sprintf("https://api.geoapify.com/v1/geocode/search?text=%s&limit=1&format=json&apiKey=%s", countryName, apiKey)
    resp, err := http.Get(url)
    if err != nil {
        return 0, 0, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != 200 {
        return 0, 0, fmt.Errorf("error fetching data for %s: %d", countryName, resp.StatusCode)
    }

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return 0, 0, err
    }

    var geocodes []Geocode
    if err := json.Unmarshal(body, &geocodes); err != nil {
        return 0, 0, err
    }

    if len(geocodes) == 0 {
        return 0, 0, fmt.Errorf("no geocoding data found for %s", countryName)
    }

    return geocodes[0].Lat, geocodes[0].Lon, nil
}