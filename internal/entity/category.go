package entity

import (
	"github.com/oaraujocesar/mpf/pkg/custom_errors"
	"github.com/oaraujocesar/mpf/pkg/entity"
)

type Category struct {
	ID   entity.ID `json:"id"`
	Name string    `json:"name"`
}

func NewCategory(name string) (*Category, error) {
	if name == "" {
		return nil, custom_errors.ErrNameRequired
	}

	if len(name) < 2 {
		return nil, custom_errors.ErrNameMinLength
	} else if len(name) > 100 {
		return nil, custom_errors.ErrNameMaxLength
	}

	return &Category{
		ID:   entity.NewID(),
		Name: name,
	}, nil
}
