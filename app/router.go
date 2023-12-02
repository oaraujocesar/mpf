package app

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/oaraujocesar/mpf/app/controllers"
	"github.com/oaraujocesar/mpf/database"
	sqlc "github.com/oaraujocesar/mpf/database/sqlc"
)

func Router() *chi.Mux {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Use(middleware.Heartbeat("/heartbeat"))
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)
	r.Use(middleware.AllowContentEncoding("application/json", "application/x-www-form-urlencoded"))
	r.Use(middleware.CleanPath)
	r.Use(middleware.RedirectSlashes)

	r.Post("/category", func(w http.ResponseWriter, r *http.Request) {
		conn, ctx := database.ConnectDB()
		queries := sqlc.New(conn)
		defer conn.Close(ctx)
		category, err := queries.CreateCategory(ctx, "test")
		if err != nil {
			w.Write([]byte(err.Error()))
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(category); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	})

	r.Route("/api/auth", func(r chi.Router) {
		r.Post("/signup", controllers.Signup)
	})

	return r
}
