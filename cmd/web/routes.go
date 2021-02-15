package main

import (
	"net/http"

	"github.com/NathanielRand/go-bnb/pkg/config"
	"github.com/NathanielRand/go-bnb/pkg/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	// Routes
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/contact", handlers.Repo.Contact)

	// Routes - Rooms
	mux.Get("/rooms/tidal", handlers.Repo.Tidal)
	mux.Get("/rooms/cliffside", handlers.Repo.Cliffside)
	mux.Get("/rooms/oceanspray", handlers.Repo.Oceanspray)
	mux.Get("/rooms/seabreeze", handlers.Repo.Seabreeze)
	mux.Get("/rooms/dawn", handlers.Repo.Dawn)
	mux.Get("/rooms/sunset", handlers.Repo.Sunset)

	// Routes - Reserve
	mux.Get("/reserve", handlers.Repo.Reservation)
	mux.Get("/reserve-cliffside", handlers.Repo.Cliffside)

	// Serve static files
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
