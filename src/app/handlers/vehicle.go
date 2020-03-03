package handlers

import (
	"fmt"
	"net/http"
)

// CreateVehicle to create new vehicle
func CreateVehicle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request:", r)
}
