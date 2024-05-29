package errors

import "errors"

var (
	ErrUserExists = errors.New("user already exists")

	ErrUserNotFound = errors.New("user not found")

	ErrUserDisabled = errors.New("user disabled")

	ErrInvalidCredentials = errors.New("invalid credentials")
)
