package custom_errors

import "errors"

var (
	ErrPasswordRequired = errors.New("password is required")
	ErrNameRequired     = errors.New("name is required")
	ErrNameMinLength    = errors.New("name must be at least 2 characters long")
	ErrNameMaxLength    = errors.New("name must be at most 100 characters long")
	ErrEmailRequired    = errors.New("email is required")
	ErrInvalidRole      = errors.New("role must be either admin or common")
	ErrUserExists       = errors.New("user already exists")
)
