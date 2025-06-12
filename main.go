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
		filtered := util.FilterCache(cache, 1)
		templ.Handler(views.TableContents(filtered))
	})

	// start server
	fmt.Println("starting server on :8080...")
	http.ListenAndServe(":8080", nil)
}
