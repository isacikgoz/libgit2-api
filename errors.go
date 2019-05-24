package git

import (
	"errors"
)

var (
	// ErrAuthenticationRequired as the name implies
	ErrAuthenticationRequired = errors.New("authentication required")
	// ErrAuthenticationType means that given credentials cannot be used for given repository url
	ErrAuthenticationType = errors.New("authentication method is not valid")
)