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

// Así van los partidos
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
	r := mux.NewRouter()

	// Rutas de la API
	r.HandleFunc("/api/matches", getMatches).Methods("GET")
	r.HandleFunc("/api/matches/{id}", getMatch).Methods("GET")
	r.HandleFunc("/api/matches", createMatch).Methods("POST")
	r.HandleFunc("/api/matches/{id}", updateMatch).Methods("PUT")
	r.HandleFunc("/api/matches/{id}", deleteMatch).Methods("DELETE")
	r.HandleFunc("/api/matches/{id}/goals", updateGoals).Methods("PATCH")
	r.HandleFunc("/api/matches/{id}/yellowcards", registerYellowCard).Methods("PATCH")
	r.HandleFunc("/api/matches/{id}/redcards", registerRedCard).Methods("PATCH")
	r.HandleFunc("/api/matches/{id}/extratime", registerExtraTime).Methods("PATCH")

	// Servir archivos estáticos
	fs := http.FileServer(http.Dir("./static"))
	r.PathPrefix("/").Handler(fs)

	// Configurar CORS
	handler := cors.Default().Handler(r)

	// Iniciar el servidor en el puerto 8080
	fmt.Println("Servidor corriendo en el puerto 8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}

// Actualizar los goles de un partido
func updateGoals(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var goal struct {
		Team   string `json:"team"`
		Player string `json:"player"`
		Minute int    `json:"minute"`
	}

	err := json.NewDecoder(r.Body).Decode(&goal)
	if err != nil {
		http.Error(w, "Error al procesar el JSON", http.StatusBadRequest)
		return
	}

	_, err = db.Exec(
		"INSERT INTO goals(match_id, team, player, minute) VALUES($1, $2, $3, $4)",
		id, goal.Team, goal.Player, goal.Minute)
	if err != nil {
		http.Error(w, "Error al registrar el gol", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Registrar tarjeta amarilla
func registerYellowCard(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var card struct {
		Team   string `json:"team"`
		Player string `json:"player"`
		Minute int    `json:"minute"`
	}

	err := json.NewDecoder(r.Body).Decode(&card)
	if err != nil {
		http.Error(w, "Error al procesar el JSON", http.StatusBadRequest)
		return
	}

	_, err = db.Exec(
		"INSERT INTO yellow_cards(match_id, team, player, minute) VALUES($1, $2, $3, $4)",
		id, card.Team, card.Player, card.Minute)
	if err != nil {
		http.Error(w, "Error al registrar la tarjeta amarilla", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Registrar tarjeta roja
func registerRedCard(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var card struct {
		Team   string `json:"team"`
		Player string `json:"player"`
		Minute int    `json:"minute"`
	}

	err := json.NewDecoder(r.Body).Decode(&card)
	if err != nil {
		http.Error(w, "Error al procesar el JSON", http.StatusBadRequest)
		return
	}

	_, err = db.Exec(
		"INSERT INTO red_cards(match_id, team, player, minute) VALUES($1, $2, $3, $4)",
		id, card.Team, card.Player, card.Minute)
	if err != nil {
		http.Error(w, "Error al registrar la tarjeta roja", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Registrar tiempo extra
func registerExtraTime(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var extraTime struct {
		ExtraTime int `json:"extra_time"`
	}

	err := json.NewDecoder(r.Body).Decode(&extraTime)
	if err != nil {
		http.Error(w, "Error al procesar el JSON", http.StatusBadRequest)
		return
	}

	_, err = db.Exec(
		"INSERT INTO extra_time(match_id, extra_time) VALUES($1, $2)",
		id, extraTime.ExtraTime)
	if err != nil {
		http.Error(w, "Error al registrar el tiempo extra", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
