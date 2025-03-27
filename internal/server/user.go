package server

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/Kaamkiya/nanoid-go"
	"github.com/go-chi/chi/v5"

	"github.com/Kaamkiya/scamalytics/internal/db"
)

func addUser(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	password := r.FormValue("password")
	webform := r.FormValue("webform") == "webform"

	if name == "" || password == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := db.AddUser(nanoid.Nanoid(nanoid.DefaultLength, nanoid.DefaultCharset), name, password)
	if err != nil {
		if webform {
			w.Write([]byte("Internal Server Error."))
		}
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Printf("%v\n", err)
		return
	}

	if !webform {
		w.WriteHeader(http.StatusCreated)
		return
	}

	http.Redirect(w, r, "http://0.0.0.0:3000", http.StatusPermanentRedirect)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	u, err := db.GetUserByID(id)
	if errors.Is(err, sql.ErrNoRows) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(u); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func loginUser(w http.ResponseWriter, r *http.Request) {
	var data map[string]string

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	r.Body.Close()

	u, err := db.GetUserByName(data["name"])
	if errors.Is(err, sql.ErrNoRows) {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	if !u.CheckPassword(data["password"]) {
		w.WriteHeader(http.StatusForbidden)
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Printf("%v\n", err)
		return
	}

	u.SID = nanoid.Default()

	err = db.SetUserSID(u.ID, u.SID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"name": u.Name,
		"id":   u.ID,
		"sid":  u.SID,
	}
	if err = json.NewEncoder(w).Encode(res); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
