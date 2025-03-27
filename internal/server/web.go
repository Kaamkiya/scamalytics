package server

import (
	"fmt"
	"html/template"
	"net/http"
)

var tmpl, err = template.ParseGlob("templates/*.gotmpl")

func webHome(w http.ResponseWriter, r *http.Request) {
	if err := tmpl.ExecuteTemplate(w, "index", map[string]string{}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func webSignup(w http.ResponseWriter, r *http.Request) {
	if err := tmpl.ExecuteTemplate(w, "signup", map[string]string{}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func webLogin(w http.ResponseWriter, r *http.Request) {
	if err := tmpl.ExecuteTemplate(w, "login", map[string]string{}); err != nil {
		fmt.Printf("%w\n", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
