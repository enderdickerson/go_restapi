package app

import (
	"log"
	"net/http"
	"restapi/src/app/handlers"
)

// App struct
type App struct{}

// Initialize application
func (a *App) Initialize() {
	setRoutes()
}

func setRoutes() {
	http.HandleFunc("/vehicle", handlers.CreateVehicle)
}

// Start application
func (a *App) Start() {
	log.Fatal(http.ListenAndServe(":8080", nil))
}
