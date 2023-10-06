package handlers

import (
	"net/http"

	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"

	"gitlab.com/rarimo/polygonid/verify-proxy/internal/service/api/responses"
	"gitlab.com/rarimo/polygonid/verify-proxy/internal/service/core"
)

func VerificationResponse(w http.ResponseWriter, r *http.Request) {
	jwz, err := VerifyProxy(r).GetJWZToken(JWTBasicClaims(r).ID)
	switch {
	case errors.Is(err, core.ErrNotVerified):
		Log(r).WithField("reason", err).Debug("No content")
		w.WriteHeader(http.StatusNoContent)
		return
	case errors.Is(err, core.ErrVerifyRequestNotFound):
		Log(r).WithField("reason", err).Debug("Not found")
		ape.RenderErr(w, problems.NotFound())
		return
	case err != nil:
		Log(r).WithError(err).Error("Failed get verify response")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, responses.NewJWZ(jwz))
}
