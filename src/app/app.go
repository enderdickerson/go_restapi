package app

import (
	"log"
	"net/http"
	h "restapi/src/app/handlers"
	m "restapi/src/app/middleware"
)

// App struct
type App struct{}

// Initialize application
func (a *App) Initialize() {
	setRoutes()
}

func setRoutes() {
	http.HandleFunc("/vehicle", m.Chain(h.CreateVehicle, m.Logging(), m.Method("POST")))
}

// Start application
func (a *App) Start() {
	log.Fatal(http.ListenAndServe(":8080", nil))
}
