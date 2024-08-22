package api

import (
	"fmt"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"gitlab.com/distributed_lab/ape"

	"github.com/rarimo/verify-proxy/internal/service/api/handlers"
	"github.com/rarimo/verify-proxy/internal/service/api/middleware"
	"github.com/rarimo/verify-proxy/internal/service/api/requests"
)

func (s *service) router() chi.Router {
	r := chi.NewRouter()

	// Basic CORS
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.ContentType("application/json"),
		ape.CtxMiddleware(
			handlers.CtxLog(s.log),
			handlers.CtxVerifyProxy(s.verifyProxy),
		),
	)
	r.Route("/integrations/verify-proxy", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Route("/public", func(r chi.Router) {
				r.Route("/verify", func(r chi.Router) {
					r.Post("/request", handlers.VerificationRequest)
					r.Get(
						fmt.Sprintf("/request/{%s}", requests.RequestIDPathParam),
						handlers.VerificationRequestData,
					)
					r.Post(
						fmt.Sprintf("/callback/{%s}", requests.RequestIDPathParam),
						handlers.VerificationCallback,
					)
					r.With(middleware.AuthMiddleware(s.jwtCfg)).
						Get(
							fmt.Sprintf("/response/{%s}", requests.RequestIDPathParam),
							handlers.VerificationResponse,
						)
				})
			})
		})
	})

	return r
}
