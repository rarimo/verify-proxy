package core

import (
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	"gitlab.com/rarimo/polygonid/verify-proxy/internal/data"
	"gitlab.com/rarimo/polygonid/verify-proxy/internal/service/api/requests"
)

func (v *verifyProxy) NewVerificationRequest() (*uuid.UUID, string, error) {
	requestID := uuid.New()
	err := v.db.VerifyRequestsQ().Insert(&data.VerifyRequest{
		ID:        requestID,
		Status:    data.VerificationStatusInitialized,
		CreatedAt: time.Now().UTC(),
	})
	if err != nil {
		return nil, "", errors.Wrap(err, "failed to insert new verify request")
	}

	jwt, err := newBasicJWT(v.jwtSecret, BasicJWTClaims{
		ID:  requestID,
		Exp: time.Now().UTC().Add(v.jwtExp).Unix(),
	})
	if err != nil {
		return nil, "", errors.Wrap(err, "failed to create jwt")
	}

	return &requestID, jwt, nil
}

func (v *verifyProxy) VerifyCallback(request *requests.VerificationCallbackRequest) error {
	verifyRequest, err := v.db.VerifyRequestsQ().WhereID(request.VerificationID).Get()
	if err != nil {
		return errors.Wrap(err, "failed to get verify request")
	}
	if verifyRequest == nil {
		return ErrVerifyRequestNotFound
	}
	if verifyRequest.Status != data.VerificationStatusInitialized {
		return ErrCallBackAlreadyProcessed
	}

	verifyRequest.CallbackData = []byte(request.JWZToken)
	verifyRequest.UpdatedAt = time.Now().UTC()
	verifyRequest.Status = data.VerificationStatusVerified
	if err = v.db.VerifyRequestsQ().Update(verifyRequest); err != nil {
		return errors.Wrap(err, "failed to update verify request")
	}

	return nil
}

func (v *verifyProxy) GetJWZToken(requestID uuid.UUID) (string, error) {
	verifyRequest, err := v.db.VerifyRequestsQ().WhereID(requestID).Get()
	if err != nil {
		return "", errors.Wrap(err, "failed to get verify request")
	}
	if verifyRequest == nil {
		return "", ErrVerifyRequestNotFound
	}
	if verifyRequest.Status != data.VerificationStatusVerified {
		return "", ErrNotVerified
	}

	return string(verifyRequest.CallbackData), nil
}
