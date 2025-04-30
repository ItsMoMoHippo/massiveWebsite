package handler

import (
	"html/template"
	"maWeb/models"
	"net/http"
)

func WharPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/message.html"))
	data := models.PageData{
		Message: ":whar:",
	}
	tmpl.Execute(w, data)
}
