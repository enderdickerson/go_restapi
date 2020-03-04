// Package handlers provide specific http endpoint operation
package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Vehicle struct for holding ymmt
type Vehicle struct {
	Year  string `json:"year"`
	Make  string `json:"make"`
	Model string `json:"model"`
	Trim  string `json:"trim"`
}

// CreateVehicle to create new vehicle
func CreateVehicle(w http.ResponseWriter, r *http.Request) {
	var v Vehicle
	err := json.NewDecoder(r.Body).Decode(&v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Vehicle: %+v", v)
	fmt.Printf("Vehicle: %+v", v)
}
