package server

import (
	"database/sql"
	"errors"
	"fmt"
	"html/template"
	"net/http"

	"github.com/Kaamkiya/scamalytics/internal/db"
	"github.com/dustin/go-humanize"
)

var tmpl *template.Template

func init() {
	tmpl, _ = template.New("root").Funcs(template.FuncMap{
		"humanizeTime": humanize.Time,
	}).ParseGlob("templates/*.gotmpl")
}

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
		fmt.Printf("%v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func webProfile(w http.ResponseWriter, r *http.Request) {
	sid, err := r.Cookie("sid")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
	}

	u, err := db.GetUserBySID(sid.Value)
	if errors.Is(err, sql.ErrNoRows) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := tmpl.ExecuteTemplate(w, "profile", u); err != nil {
		fmt.Printf("%v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func webLessons(w http.ResponseWriter, r *http.Request) {
	sid, err := r.Cookie("sid")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
	}

	u, err := db.GetUserBySID(sid.Value)
	if errors.Is(err, sql.ErrNoRows) {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := tmpl.ExecuteTemplate(w, "lessons", u); err != nil {
		fmt.Printf("%v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
