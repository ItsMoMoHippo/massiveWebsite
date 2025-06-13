package main

import (
	"fmt"
	"log"
	"maWeb/db"
	"maWeb/dbquery"
	"maWeb/model"
	"maWeb/util"
	"maWeb/views"
	"net/http"
	"strconv"

	"github.com/a-h/templ"
)

var cache *[]model.TrackTime

// main function
func main() {
	// db startup
	db, err := db.LInit()
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	defer db.Close()

	// db tests
	fmt.Println("pinging db...")
	db.Ping()
	fmt.Println("db ponged")
	if err = dbquery.PrintStar(db); err != nil {
		log.Fatal(err)
	}

	// init cache
	tempCache, err := dbquery.GetStar(db)
	if err != nil {
		log.Fatal(err)
	}
	cache = &tempCache
	util.PrintCache(*cache)

	// refresh cache
	// TODO: put in a go routine
	refreshedCach, err := util.RefreshCache(db)
	if err != nil {
		log.Fatal(err)
	}
	cache = &refreshedCach

	// handles
	http.Handle("/", templ.Handler(views.Page("MassiveAttackWeb", "Massive Attack", cache)))

	// TODO: do endpoint
	// returns filtered times
	http.HandleFunc("/filter", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			fmt.Println("error")
			return
		}
		fmt.Println("data stage 1 accepted")

		clubIDstr := r.FormValue("club_id")
		if clubIDstr == "" {
			http.Error(w, "Missing club_id", http.StatusBadRequest)
			fmt.Println("missing club_id")
			return
		}
		fmt.Println("data satge 2 accepted")

		clubID, err := strconv.Atoi(clubIDstr)
		if err != nil {
			http.Error(w, "Invalid Club Id", http.StatusBadRequest)
			fmt.Println("data was not a number")
			return
		}
		fmt.Println("data stage 3 accepted")

		filtered := util.FilterCache(cache, clubID)
		util.PrintCache(*filtered)
		fmt.Println("making data")
		templ.Handler(views.TableContents(filtered)).ServeHTTP(w, r)
		fmt.Println("sending data")
	})

	// start server
	fmt.Println("starting server on :8080...")
	http.ListenAndServe(":8080", nil)
}
