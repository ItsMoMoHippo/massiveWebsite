package main

import (
	"fmt"
	"maWeb/handler"
	"net/http"

	_ "github.com/a-h/templ"
	_ "github.com/lib/pq"
)

// main function
func main() {
	http.HandleFunc("/", handler.HomePage)
	http.HandleFunc("/wharswap", handler.WharPage)
	http.HandleFunc("/setClubFoo", handler.SetClubFoo)
	http.HandleFunc("/changeTimes", handler.DifferentList)

	fmt.Println("Starting sevrer on :8080...")
	http.ListenAndServe(":8080", nil)
}
