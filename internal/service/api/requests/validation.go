package requests

import (
	"github.com/google/uuid"
	"github.com/iden3/go-jwz"
	"github.com/pkg/errors"
)

func MustBeValidJWZToken(src interface{}) error {
	tokenRaw, ok := src.(string)
	if !ok {
		return errors.New("it is not a JWZ token")
	}

	_, err := jwz.Parse(tokenRaw)
	if err != nil {
		return errors.New("it is not a valid JWZ token")
	}

	return nil
}

func MustBeValidUUID(src interface{}) error {
	uuidRaw, ok := src.(string)
	if !ok {
		return errors.New("it is not a string")
	}

	_, err := uuid.Parse(uuidRaw)
	if err != nil {
		return errors.New("it is not a valid uuid")
	}

	return nil
}
