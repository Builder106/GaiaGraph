package models

type ClimateData struct {
    ID        int     `json:"id"`
    Metric    string  `json:"metric"`
    Value     float64 `json:"value"`
    Timestamp string  `json:"timestamp"`
}