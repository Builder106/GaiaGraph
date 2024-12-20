package externalapis

import (
    "encoding/json"
    "fmt"
    "io"
    "log"
    "net/http"

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
	// Additional fields can be added here if needed
}

type AddressComponents struct {
	City     string `json:"city"`
	State    string `json:"state"`
	Country  string `json:"country"`
	Postcode string `json:"postcode"`
	// Additional components can be added here if needed
}


func LoadEnv() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
}

func GetCountryGeocode(countryName, apiKey string) (float64, float64, error) {
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

	var geocodeResponse GeocodeResponse
	if err := json.Unmarshal(body, &geocodeResponse); err != nil {
		 return 0, 0, err
	}

	if len(geocodeResponse.Results) == 0 {
		 return 0, 0, fmt.Errorf("no geocoding data found for %s", countryName)
	}

	result := geocodeResponse.Results[0]
	// Validate that the result corresponds to the intended country
	if result.Components.Country != countryName {
		 return 0, 0, fmt.Errorf("geocoding result does not match the country name: %s", countryName)
	}

	return result.Lat, result.Lon, nil
}