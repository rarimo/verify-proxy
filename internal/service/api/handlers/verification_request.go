package handlers

import (
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"

	"gitlab.com/rarimo/polygonid/verify-proxy/internal/service/api/responses"
)

func VerificationRequest(w http.ResponseWriter, r *http.Request) {
	requestID, jwt, err := VerifyProxy(r).NewVerificationRequest()
	if err != nil {
		Log(r).WithError(err).Debug("Internal error")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, responses.NewVerificationID(requestID, jwt))
}
