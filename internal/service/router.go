package service

import (
	"github.com/go-chi/chi"
	"github.com/kish1n/shortlink/internal/service/handlers"
	"github.com/kish1n/shortlink/internal/service/helpers"
	"github.com/kish1n/shortlink/internal/service/requests"
	"gitlab.com/distributed_lab/ape"
	"log"
	"net/http"
)

func (s *service) router() (chi.Router, error) {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			helpers.CtxLog(s.log),
		),
	)

	r.Route("/integrations/shortlink", func(r chi.Router) {
		r.Get("/db", requests.DBHandler)
		r.Post("/add", handlers.GetShort)
		r.Get("/{shortened}", requests.RedirectHandler)
	})

	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", r)

	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
		return r, err
	}

	return r, nil
}
