package responses

import (
	//"github.com/rarimo/verify-proxy/resources"

	"github.com/google/uuid"

	"github.com/rarimo/verify-proxy/resources"
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
