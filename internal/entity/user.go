package entity

import (
	"github.com/oaraujocesar/mpf/pkg/custom_errors"
	"github.com/oaraujocesar/mpf/pkg/entity"
	"golang.org/x/crypto/bcrypt"
)

type Role string

const (
	Admin  Role = "admin"
	Common Role = "common"
)

type User struct {
	ID       entity.ID `json:"id"`
	Name     string    `json:"name"`
	Avatar   string    `json:"avatar"`
	Email    string    `json:"email"`
	Password string    `json:"-"`
	Role     Role      `json:"role"`
	Entries  []Entry   `json:"entries"`
}

func NewUser(name, avatar, email, password string, role Role) (*User, error) {
	if name == "" {
		return nil, custom_errors.ErrNameRequired
	}

	if len(name) < 2 {
		return nil, custom_errors.ErrNameMinLength
	} else if len(name) > 100 {
		return nil, custom_errors.ErrNameMaxLength
	}

	if email == "" {
		return nil, custom_errors.ErrEmailRequired
	}

	if password == "" {
		return nil, custom_errors.ErrPasswordRequired
	} else if len(password) < 10 {
		return nil, custom_errors.ErrPasswordMinLength
	}

	if role != Admin && role != Common {
		return nil, custom_errors.ErrInvalidRole
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:       entity.NewID(),
		Name:     name,
		Avatar:   avatar,
		Email:    email,
		Password: string(hash),
		Role:     role,
	}, nil
}

func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))

	return err == nil
}
