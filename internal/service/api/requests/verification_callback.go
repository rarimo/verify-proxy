package requests

import (
	"io"
	"net/http"

	"github.com/go-chi/chi"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type VerificationCallbackRequest struct {
	VerificationID uuid.UUID
	JWZToken       string
}

type verificationCallbackRequest struct {
	VerificationID string
	JWZToken       string
}

func NewVerificationCallbackRequest(r *http.Request) (*VerificationCallbackRequest, error) {
	tokenRaw, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, errors.New("is not a valid request token")
	}

	requestRaw := verificationCallbackRequest{
		JWZToken:       string(tokenRaw),
		VerificationID: chi.URLParam(r, RequestIDPathParam),
	}

	if err := requestRaw.validate(); err != nil {
		return nil, err
	}

	return requestRaw.parse(), nil
}

func (req *verificationCallbackRequest) parse() *VerificationCallbackRequest {
	return &VerificationCallbackRequest{
		VerificationID: uuid.MustParse(req.VerificationID),
		JWZToken:       req.JWZToken,
	}
}

func (req *verificationCallbackRequest) validate() error {
	return validation.Errors{
		"path/{request-id}": validation.Validate(
			req.VerificationID, validation.Required, validation.By(MustBeValidUUID),
		),
		"body": validation.Validate(
			req.JWZToken, validation.Required, validation.By(MustBeValidJWZToken),
		),
	}.Filter()
}
