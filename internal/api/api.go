package api

import (
    "encoding/json"
    "net/http"
    "GaiaGraph/internal/db"
    "GaiaGraph/internal/models"
    "io"
)

func GetClimateData(w http.ResponseWriter, r *http.Request) {
    rows, err := db.DB.Query("SELECT id, metric, value, timestamp FROM climate_data")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var data []models.ClimateData
    for rows.Next() {
        var d models.ClimateData
        if err := rows.Scan(&d.ID, &d.Metric, &d.Value, &d.Timestamp); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        data = append(data, d)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(data)
}

type ClimateData struct {
    Metric    string  `json:"metric"`
    Value     float64 `json:"value"`
    Timestamp string  `json:"timestamp"`
}

func GetCombinedData(w http.ResponseWriter, r *http.Request) {
    // Fetch data from the first external API
    resp1, err := http.Get("https://api.external1.com/climate-data")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer resp1.Body.Close()

    body1, err := io.ReadAll(resp1.Body)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    var data1 []ClimateData
    if err := json.Unmarshal(body1, &data1); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Fetch data from the second external API
    resp2, err := http.Get("https://api.external2.com/climate-data")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer resp2.Body.Close()

    body2, err := io.ReadAll(resp2.Body)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    var data2 []ClimateData
    if err := json.Unmarshal(body2, &data2); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Combine the data from both APIs
    combinedData := append(data1, data2...)

    // Return the combined data as JSON
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(combinedData)
}