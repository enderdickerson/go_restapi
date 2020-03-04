// Package handlers provide specific http endpoint operation
package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"restapi/src/app/data"
	"restapi/src/app/models"

	// Import for postgres dialect
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// CreateVehicle to create new vehicle
func CreateVehicle(w http.ResponseWriter, r *http.Request) {
	var v models.Vehicle
	err := json.NewDecoder(r.Body).Decode(&v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data.GetDB().Create(&v)

	fmt.Fprintf(w, "Vehicle: %+v", v)
	fmt.Printf("Vehicle: %+v", v)
}
