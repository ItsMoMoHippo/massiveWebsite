package handler

import (
	"html/template"
	"maWeb/models"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	data := models.PageData{
		Title:   "MA Website",
		Header:  "Welcome from Go",
		Message: ":michiraisedeyebrows:",
		Times:   []string{"00:00", "00:01", "00:02"},
	}
	tmpl.Execute(w, data)
}
