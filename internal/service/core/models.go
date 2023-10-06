package core

import "gitlab.com/distributed_lab/logan/v3/errors"

var (
	ErrCallBackAlreadyProcessed = errors.New("callback already processed")
	ErrNotVerified              = errors.New("not verified")
	ErrVerifyRequestNotFound    = errors.New("verify request not found")
)
