package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	var err error

	connStr := "user=postgres password=taylorswift dbname=liga_tracker sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
}

func getTeams(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT * FROM teams")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var teams []map[string]interface{}

	for rows.Next() {
		var id int
		var name, logoUrl string
		if err := rows.Scan(&id, &name, &logoUrl); err != nil {
			log.Fatal(err)
		}
		teams = append(teams, map[string]interface{}{
			"id":      id,
			"name":    name,
			"logoUrl": logoUrl,
		})
	}

	json.NewEncoder(w).Encode(teams)
}

func serveHTML(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w, r, "LaLigaTracker.html")
}

func main() {

	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/teams", getTeams)

	http.HandleFunc("/", serveHTML)

	fmt.Println("Servidor corriendo en puerto 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
