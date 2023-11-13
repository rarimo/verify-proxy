package responses

import "github.com/rarimo/verify-proxy/resources"

func NewJWZ(jwz string) *resources.JwzResponse {
	return &resources.JwzResponse{
		Data: resources.Jwz{
			Key: resources.Key{
				Type: resources.JWZ,
			},
			Attributes: resources.JwzAttributes{
				Jwz: jwz,
			},
		},
	}
}
