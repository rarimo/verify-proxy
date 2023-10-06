package core

import (
	"github.com/fatih/structs"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

const (
	IDClaimName  = "id"
	ExpClaimName = "exp"
)

type BasicJWTClaims struct {
	ID  uuid.UUID `structs:"id"`
	Exp int64     `structs:"exp"`
}

type AccessLevel string

func newBasicJWT(
	secretKey []byte,
	claims BasicJWTClaims,
) (string, error) {
	signedToken, err := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims(structs.Map(claims)),
	).SignedString(secretKey)
	if err != nil {
		return "", errors.Wrap(err, "failed to sign token")
	}

	return signedToken, nil
}
