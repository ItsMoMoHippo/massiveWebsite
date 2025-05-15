package main

import (
	"fmt"
	"log"
	"maWeb/db"
	"maWeb/dbquery"
	"maWeb/htmlstuff"
	"net/http"

	"github.com/a-h/templ"
)

type Filter struct {
	Name string
}

var filter = Filter{Name: "Joe"}

// MAIN FUNCTION
func main() {
	db, err := db.LInit()
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	defer db.Close()

	http.Handle("/", templ.Handler(htmlstuff.Page("MassiveAttackWeb", "Massive Attack", "Hello from templ")))
	http.Handle("/msgswap", templ.Handler(htmlstuff.Message("hello?")))
	http.Handle("/listswap", templ.Handler(htmlstuff.ListSwap([]string{"00.03", "00.04"})))

	fmt.Println("pinging db...")
	db.Ping()
	fmt.Println("db ponged")

	if err := dbquery.PrintAll(db); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/tracktimes", func(w http.ResponseWriter, r *http.Request) {
		data, err := dbquery.GetTrackTimes(db)
		if err != nil {
			http.Error(w, "query error", 500)
			return
		}

		w.Header().Set("Content-Type", "text/html")
		htmlstuff.TrackTimesTable(data).Render(r.Context(), w)
	})
	http.HandleFunc("/changefilter", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("before is: ", filter.Name)
		filter.Name = "Bob"
		fmt.Println("Filter changed to:", filter.Name)
		w.WriteHeader(http.StatusOK)
	})

	fmt.Println("starting server on :8080...")
	http.ListenAndServe(":8080", nil)
}
