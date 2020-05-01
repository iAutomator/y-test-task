package main

import (
	"html/template"
	"net/http"
	"y-test-task/visits"
)

func main() {
	view := template.Must(setupView())
	http.Handle("/", visits.NewController(view))
	http.ListenAndServe(":9999", nil)
}

func setupView() (*template.Template, error) {
	return template.ParseFiles("resources/index.html")
}
