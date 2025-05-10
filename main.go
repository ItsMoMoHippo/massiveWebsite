package main

import (
	"fmt"
	"log"
	"maWeb/db"
	"maWeb/htmlstuff"
	"net/http"

	"github.com/a-h/templ"
	_ "github.com/lib/pq"
)

// main function
func main() {
	db, err := db.LInit()
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	http.Handle("/", templ.Handler(htmlstuff.Page("MassiveAttackWeb", "MassiveAttack", "hello from templ", []string{"00.01", "00.02"})))
	http.Handle("/msgswap", templ.Handler(htmlstuff.Message("hello?")))
	http.Handle("/listswap", templ.Handler(htmlstuff.ListSwap([]string{"00.03", "00.04"})))

	db.Ping()
	fmt.Println("db ponged")
	fmt.Println("Starting sevrer on :8080...")
	http.ListenAndServe(":8080", nil)
}
