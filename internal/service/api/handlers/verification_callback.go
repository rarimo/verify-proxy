package handlers

import (
	"net/http"

	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"

	"github.com/rarimo/verify-proxy/internal/service/api/requests"
	"github.com/rarimo/verify-proxy/internal/service/core"
)

func VerificationCallback(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewVerificationCallbackRequest(r)
	if err != nil {
		Log(r).WithField("reason", err).Debug("Bad request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	err = VerifyProxy(r).VerifyCallback(req)
	switch {
	case errors.Is(err, core.ErrVerifyRequestNotFound):
		Log(r).WithField("reason", err).Debug("Not found")
		ape.RenderErr(w, problems.NotFound())
		return
	case errors.Is(err, core.ErrCallBackAlreadyProcessed):
		Log(r).WithField("reason", err).Debug("Conflict")
		ape.RenderErr(w, problems.Conflict())
		return
	case err != nil:
		Log(r).WithError(err).Error("Failed process callback")
		ape.RenderErr(w, problems.InternalError())
		return
	}
}
