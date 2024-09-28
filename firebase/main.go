package main

import (
	"encoding/json"
	"net/http"
)

// Updated Config struct to include all necessary fields
type Config struct {
	APIKey             string `json:"apiKey"`
	AuthDomain         string `json:"authDomain"`
	DatabaseURL        string `json:"databaseURL"`
	ProjectID          string `json:"projectId"`
	StorageBucket      string `json:"storageBucket"`
	MessagingSenderID  string `json:"messagingSenderId"`
	AppID              string `json:"appId"`
	MeasurementID      string `json:"measurementId"`
}

func getAPIKey(w http.ResponseWriter, r *http.Request) {
	config := Config{
		APIKey:             "AIzaSyA-Dl0UYvx2na0oG0mXPPSGxRPxe7-uMfE",
		AuthDomain:         "bus-tracking-2b76e.firebaseapp.com",
		DatabaseURL:        "https://bus-tracking-2b76e-default-rtdb.asia-southeast1.firebasedatabase.app",
		ProjectID:          "bus-tracking-2b76e",
		StorageBucket:      "bus-tracking-2b76e.appspot.com",
		MessagingSenderID:  "964312599473",
		AppID:              "1:964312599473:web:82dc639eaf2d6451a48b90",
		MeasurementID:      "G-HDBKHSTKML",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(config)
}

func main() {
	http.HandleFunc("/get-api-key", getAPIKey)
	http.ListenAndServe(":8080", nil) // Server will run on port 8080
}
