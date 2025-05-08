package main

import (
	"fmt"
	"maWeb/htmlstuff"
	"net/http"

	"github.com/a-h/templ"
	_ "github.com/lib/pq"
)

// main function
func main() {
	homePage := htmlstuff.Page("MassiveAttackWeb", "MassiveAttack", "hello from templ")

	http.Handle("/", templ.Handler(homePage))
	http.Handle("/msgswap", templ.Handler(htmlstuff.Message("hello?")))

	fmt.Println("Starting sevrer on :8080...")
	http.ListenAndServe(":8080", nil)
}
