package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/lib/pq"
)

// connects to the postgreSQL DB
func connectDB() *sql.DB {
	//add thingy to grab from env file or smt
	connStr := "postgres creds"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db
}

// retrieves times from the db
func getTimes(db *sql.DB) ([]string, error) {
	rows, err := db.Query("SELECT foo FROM bar")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var times []string
	for rows.Next() {
		var time string
		if err := rows.Scan(&time); err != nil {
			return nil, err
		}
		times = append(times, time)
	}
	return times, rows.Err()
}

type PageData struct {
	Title   string
	Header  string
	Message string
}

// struct for holding times
type RaceTimes struct {
	Times []string
}

// loads the home page
func handler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	data := PageData{
		Title:   "MA Website",
		Header:  "Welcome from Go",
		Message: ":michiraisedeyebrows:",
	}
	tmpl.Execute(w, data)
}

// main function
func main() {
	http.HandleFunc("/", handler)

	fmt.Println("Starting sevrer on :8080...")
	http.ListenAndServe(":8080", nil)
}
