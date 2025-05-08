package handler

import (
	"html/template"
	"maWeb/models"
	"net/http"
)

func DifferentList(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/list.html"))
	data := models.TimesList{
		Times: []string{"00:03", "00:40"},
	}
	tmpl.Execute(w, data)
}
