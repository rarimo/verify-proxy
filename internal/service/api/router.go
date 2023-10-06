package api

import (
	"fmt"

	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"

	"gitlab.com/rarimo/polygonid/verify-proxy/internal/service/api/handlers"
	"gitlab.com/rarimo/polygonid/verify-proxy/internal/service/api/middleware"
	"gitlab.com/rarimo/polygonid/verify-proxy/internal/service/api/requests"
)

func (s *service) router() chi.Router {
	r := chi.NewRouter()

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
					r.Get("/request", handlers.VerificationRequest)
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
