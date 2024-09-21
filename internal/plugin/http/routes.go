package http

import (
	i "Odnostranishka/internal/plugin/interface"
	"Odnostranishka/pkg/http/router"
	"github.com/go-chi/chi/v5"
)

func Init(r chi.Router, uc i.UC) {
	//r.Route("/", func(r chi.Router) {
	r.Post("/order", router.FromBody(uc.Create))
	//})
}
