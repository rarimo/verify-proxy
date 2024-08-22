package handlers

import (
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"

	"github.com/rarimo/verify-proxy/internal/service/api/responses"
)

func VerificationRequestData(w http.ResponseWriter, r *http.Request) {
	requestData, err := VerifyProxy(r).GetVerificationRequest(r)
	if err != nil {
		Log(r).WithError(err).Debug("Internal error")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, responses.NewVerificationData(&requestData))
}
