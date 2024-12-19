package api

import (
    "encoding/json"
    "net/http"
    "GaiaGraph/internal/db"
    "GaiaGraph/internal/models"
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