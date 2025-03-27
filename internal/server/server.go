package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/Kaamkiya/scamalytics/internal/db"
)

func Run(addr string) {
	if err := db.Init(); err != nil {
		panic(err)
	}
	defer db.Close()

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/api/adduser", addUser)
	r.Get("/api/user/{id}", getUser)
	r.Post("/api/login", loginUser)

	r.Get("/", webHome)
	r.Get("/signup", webSignup)
	r.Get("/login", webLogin)
	r.Get("/user/{name}", webProfile)

	fs := http.FileServer(http.Dir("static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))

	http.ListenAndServe(addr, r)
}
