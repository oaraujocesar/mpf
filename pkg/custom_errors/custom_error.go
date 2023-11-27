package custom_errors

import "errors"

var (
	ErrPasswordRequired  = errors.New("password is required")
	ErrPasswordMinLength = errors.New("password must be at least 10 characters long")
	ErrNameRequired      = errors.New("name is required")
	ErrNameMinLength     = errors.New("name must be at least 2 characters long")
	ErrNameMaxLength     = errors.New("name must be at most 100 characters long")
	ErrTitleRequired     = errors.New("title is required")
	ErrTitleMinLength    = errors.New("title must be at least 2 characters long")
	ErrTitleMaxLength    = errors.New("title must be at most 100 characters long")
	ErrAmountRequired    = errors.New("amount is required")
	ErrInvalidType       = errors.New("type must be either 'in' or 'out'")
	ErrCategoryRequired  = errors.New("category is required")
	ErrEmailRequired     = errors.New("email is required")
	ErrInvalidRole       = errors.New("role must be either admin or common")
	ErrUserExists        = errors.New("user already exists")
	ErrUserIDRequired    = errors.New("user ID is required")
)
