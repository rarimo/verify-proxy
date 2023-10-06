package responses

import (
	//"gitlab.com/rarimo/polygonid/verify-proxy/resources"

	"github.com/google/uuid"

	"gitlab.com/rarimo/polygonid/verify-proxy/resources"
)

func NewVerificationID(requestID *uuid.UUID, jwt string) *resources.VerifyIdResponse {
	return &resources.VerifyIdResponse{
		Data: resources.VerifyId{
			Key: resources.Key{
				Type: resources.VERIFICATION_ID,
			},
			Attributes: resources.VerifyIdAttributes{
				VerificationId: requestID.String(),
				Jwt:            jwt,
			},
		},
	}
}
