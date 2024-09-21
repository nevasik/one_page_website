package server

import (
	"Odnostranishka/internal/plugin/http"
	"Odnostranishka/internal/plugin/uc"
	"Odnostranishka/internal/server/shared/middleware"
	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (s *Server) init() error {
	var (
		r *chi.Mux
	)

	uc := uc.New()

	mw := middleware.New()

	r = chi.NewRouter()
	r.Use(chimw.Logger)
	r.Use(chimw.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"*"},
	}))

	txR := r.Route("/api", func(r chi.Router) {
		r.Use(mw.Verify)
	})

	// server
	s.httpServer.Handler = r

	// routes
	http.Init(txR, uc)

	return nil
}
