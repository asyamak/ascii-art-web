package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func templates(w http.ResponseWriter, parseFile string, data interface{}) {
	html, err := template.ParseFiles(parseFile)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "500InternalServerError.html", http.StatusInternalServerError)
		return
	}
	err = html.Execute(w, data)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "500InternalServerError.html", http.StatusInternalServerError)
		return
	}
}
