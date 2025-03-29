package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

var db *sql.DB

// Estructura de los partidos
type Match struct {
	ID        int    `json:"id"`
	HomeTeam  string `json:"home_team"`
	AwayTeam  string `json:"away_team"`
	MatchDate string `json:"match_date"`
}

func init() {
	var err error
	// Conexión a la base de datos
	connStr := "user=postgres password=taylorswift dbname=liga_tracker sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
}

// Obtener todos los partidos
func getMatches(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT * FROM matches")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var matches []map[string]interface{}
	for rows.Next() {
		var id int
		var home_team, away_team, match_date string
		if err := rows.Scan(&id, &home_team, &away_team, &match_date); err != nil {
			log.Fatal(err)
		}
		matches = append(matches, map[string]interface{}{
			"id":         id,
			"home_team":  home_team,
			"away_team":  away_team,
			"match_date": match_date,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(matches)
}

// Obtener un partido por ID
func getMatch(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var match Match
	err := db.QueryRow("SELECT * FROM matches WHERE id = $1", id).Scan(&match.ID, &match.HomeTeam, &match.AwayTeam, &match.MatchDate)
	if err != nil {
		http.Error(w, "No se encontró el partido", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(match)
}

// Crear un nuevo partido
func createMatch(w http.ResponseWriter, r *http.Request) {
	var match Match
	err := json.NewDecoder(r.Body).Decode(&match)
	if err != nil {
		http.Error(w, "Error al procesar el JSON", http.StatusBadRequest)
		return
	}

	err = db.QueryRow(
		"INSERT INTO matches(home_team, away_team, match_date) VALUES($1, $2, $3) RETURNING id",
		match.HomeTeam, match.AwayTeam, match.MatchDate).Scan(&match.ID)
	if err != nil {
		http.Error(w, "Error al crear el partido", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(match)
}

// Actualizar un partido
func updateMatch(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var match Match
	err := json.NewDecoder(r.Body).Decode(&match)
	if err != nil {
		http.Error(w, "Error al procesar el JSON", http.StatusBadRequest)
		return
	}

	_, err = db.Exec(
		"UPDATE matches SET home_team = $1, away_team = $2, match_date = $3 WHERE id = $4",
		match.HomeTeam, match.AwayTeam, match.MatchDate, id)
	if err != nil {
		http.Error(w, "Error al actualizar el partido", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(match)
}

// Eliminar un partido
func deleteMatch(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	_, err := db.Exec("DELETE FROM matches WHERE id = $1", id)
	if err != nil {
		http.Error(w, "Error al eliminar el partido", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func main() {
	//manejar las rutas
	r := mux.NewRouter()

	// Rutas para manejar los partidos
	r.HandleFunc("/api/matches", getMatches).Methods("GET")
	r.HandleFunc("/api/matches/{id}", getMatch).Methods("GET")
	r.HandleFunc("/api/matches", createMatch).Methods("POST")
	r.HandleFunc("/api/matches/{id}", updateMatch).Methods("PUT")
	r.HandleFunc("/api/matches/{id}", deleteMatch).Methods("DELETE")

	// Configurar CORS
	handler := cors.Default().Handler(r)

	// Iniciar el servidor en el puerto 8080
	fmt.Println("Servidor corriendo en el puerto 8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
