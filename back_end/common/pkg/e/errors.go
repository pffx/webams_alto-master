package e

import "errors"

var (
	ErrTokenAuthFailed = errors.New(GetMessage(ERROR_AUTH_CHECK_TOKEN_FAIL))
	ErrTokenTimeOut    = errors.New(GetMessage(ERROR_AUTH_CHECK_TOKEN_TIMEOUT))
	ErrTokenInvalid    = errors.New(GetMessage(ERROR_AUTH))
	ErrDBInvalid       = errors.New("DB is not ready")
)
