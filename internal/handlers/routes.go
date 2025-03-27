package handlers

import (
	"fmt"
	"net/http"
)

// Ruta principal
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Bienvenido a La Liga Tracker API")
}

// Ruta de prueba para equipos
func TeamsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Aquí irán los equipos de La Liga")
}

// Configurar las rutas
func RegisterRoutes() {
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/teams", TeamsHandler)
}
