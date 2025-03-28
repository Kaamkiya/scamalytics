package server

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"github.com/Kaamkiya/scamalytics/internal/db"
	"github.com/dustin/go-humanize"
	"github.com/go-chi/chi/v5"
)

var tmpl *template.Template

func init() {
	tmpl, _ = template.New("root").Funcs(template.FuncMap{
		"humanizeTime": humanize.Time,
	}).ParseGlob("templates/*.gotmpl")
}

func writeTemplate(w http.ResponseWriter, name string, data any, status int) error {
	err := tmpl.ExecuteTemplate(w, name, data)

	w.WriteHeader(status)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Printf("%v\n", err)
	}

	return err
}

func webHome(w http.ResponseWriter, r *http.Request) {
	writeTemplate(w, "index", map[any]any{}, http.StatusInternalServerError)
}

func webSignup(w http.ResponseWriter, r *http.Request) {
	writeTemplate(w, "signup", map[any]any{}, http.StatusInternalServerError)
}

func webLogin(w http.ResponseWriter, r *http.Request) {
	writeTemplate(w, "login", map[any]any{}, http.StatusInternalServerError)
}

func webProfile(w http.ResponseWriter, r *http.Request) {
	sid, err := r.Cookie("sid")
	if err != nil {
		writeTemplate(
			w,
			"error",
			WebError{http.StatusUnauthorized},
			http.StatusUnauthorized,
		)
		return
	}

	u, err := db.GetUserBySID(sid.Value)
	if errors.Is(err, sql.ErrNoRows) {
		writeTemplate(w, "error", WebError{http.StatusNotFound}, http.StatusNotFound)
		return
	}

	if err != nil {
		writeTemplate(w, "error", WebError{http.StatusInternalServerError}, http.StatusInternalServerError)
		return
	}

	writeTemplate(w, "profile", u, http.StatusOK)
}

func webCourses(w http.ResponseWriter, r *http.Request) {
	sid, err := r.Cookie("sid")
	if err != nil {
		writeTemplate(w, "error", WebError{http.StatusUnauthorized}, http.StatusUnauthorized)
		return
	}

	_, err = db.GetUserBySID(sid.Value)
	if errors.Is(err, sql.ErrNoRows) {
		writeTemplate(w, "error", WebError{http.StatusForbidden}, http.StatusForbidden)
		return
	}

	if err != nil {
		writeTemplate(w, "error", WebError{http.StatusInternalServerError}, http.StatusInternalServerError)
		return
	}

	writeTemplate(w, "courses", courses, http.StatusInternalServerError)
}

func webCourse(w http.ResponseWriter, r *http.Request) {
	sid, err := r.Cookie("sid")
	if err != nil {
		writeTemplate(w, "error", WebError{http.StatusUnauthorized}, http.StatusUnauthorized)
		return
	}

	_, err = db.GetUserBySID(sid.Value)
	if errors.Is(err, sql.ErrNoRows) {
		writeTemplate(w, "error", WebError{http.StatusInternalServerError}, http.StatusInternalServerError)
		return
	}

	if err != nil {
		writeTemplate(w, "error", WebError{http.StatusInternalServerError}, http.StatusInternalServerError)
		return
	}

	slug := chi.URLParam(r, "slug")
	c, ok := courses[slug]
	if !ok {
		writeTemplate(w, "error", WebError{http.StatusNotFound}, http.StatusNotFound)
		return
	}

	writeTemplate(w, "course", c, http.StatusInternalServerError)
}

func webArticle(w http.ResponseWriter, r *http.Request) {
	sid, err := r.Cookie("sid")
	if err != nil {
		writeTemplate(w, "error", WebError{http.StatusUnauthorized}, http.StatusUnauthorized)
		return
	}

	_, err = db.GetUserBySID(sid.Value)
	if errors.Is(err, sql.ErrNoRows) {
		writeTemplate(w, "error", WebError{http.StatusForbidden}, http.StatusForbidden)
		return
	}

	if err != nil {
		writeTemplate(w, "error", WebError{http.StatusInternalServerError}, http.StatusInternalServerError)
		return
	}

	slug := chi.URLParam(r, "slug")
	sidx := chi.URLParam(r, "idx")
	idx, err := strconv.Atoi(sidx)
	if err != nil {
		writeTemplate(w, "error", WebError{http.StatusBadRequest}, http.StatusBadRequest)
		return
	}

	c, ok := courses[slug]
	if !ok {
		writeTemplate(w, "error", WebError{http.StatusNotFound}, http.StatusNotFound)
		return
	}

	writeTemplate(w, "article", c.Articles[idx], http.StatusOK)
}
