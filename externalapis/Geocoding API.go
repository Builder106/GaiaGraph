package externalapis

import (
    "encoding/json"
    "fmt"
    "io"
    "log"
    "net/http"
    "os"
    "strings"

    "github.com/joho/godotenv"
)

type GeocodeResponse struct {
    Results []GeocodeResult `json:"results"`
}

type GeocodeResult struct {
    Lat        float64           `json:"lat"`
    Lon        float64           `json:"lon"`
    Formatted  string            `json:"formatted"`
    Components AddressComponents `json:"components"`
}

type AddressComponents struct {
    City     string `json:"city"`
    State    string `json:"state"`
    Country  string `json:"country"`
    Postcode string `json:"postcode"`
}

var countryNameMapping = map[string]string{
    "Czech Republic": "Czechia",
    "Korea, North": "North Korea",
    "Korea, South": "South Korea",
    "Palestine": "Palestinian Territories",
    "Sao Tome and Principe": "São Tomé and Príncipe",
    "Timor-Leste": "East Timor",
    "Cabo Verde": "Cape Verde",
    "Cote d'Ivoire": "Côte d'Ivoire",
    "Congo, Republic of the": "Congo",
    "Congo, Democratic Republic of the": "Democratic Republic of the Congo",
}

func LoadEnv() {
    wd, err := os.Getwd()
    if err != nil {
        log.Fatalf("Error getting working directory: %v", err)
    }
    log.Printf("Current working directory: %s", wd)

    err = godotenv.Load("../.env")
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }
    log.Println(".env file loaded successfully")
}

func GetCountryGeocode(countryName, apiKey string) (float64, float64, error) {
    if altName, exists := countryNameMapping[countryName]; exists {
        countryName = altName
    }

    url := fmt.Sprintf("https://api.geoapify.com/v1/geocode/search?text=%s&limit=1&type=country&format=json&apiKey=%s", countryName, apiKey)
    resp, err := http.Get(url)
    if err != nil {
        return 0, 0, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return 0, 0, fmt.Errorf("error fetching data for %s: %d", countryName, resp.StatusCode)
    }

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return 0, 0, err
    }

    var geoapifyResponse struct {
        Results []struct {
            Lat       float64 `json:"lat"`
            Lon       float64 `json:"lon"`
            Formatted string  `json:"formatted"`
            Components struct {
                Country string `json:"country"`
            } `json:"components"`
        } `json:"results"`
    }
    if err := json.Unmarshal(body, &geoapifyResponse); err != nil {
        return 0, 0, err
    }

    if len(geoapifyResponse.Results) == 0 {
        return 0, 0, fmt.Errorf("no geocoding data found for %s", countryName)
    }

    result := geoapifyResponse.Results[0]
    log.Printf("Geocoding result for %s: %+v", countryName, result)

    if !strings.Contains(result.Formatted, countryName) {
        return 0, 0, fmt.Errorf("geocoding result does not match the country name: %s", countryName)
    }

    return result.Lat, result.Lon, nil
}