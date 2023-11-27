package entity

import (
	"github.com/google/uuid"
	"github.com/oaraujocesar/mpf/pkg/custom_errors"
	"github.com/oaraujocesar/mpf/pkg/entity"
)

type Type string

const (
	In  Type = "in"
	Out Type = "out"
)

type Entry struct {
	Title      string    `json:"title"`
	Amount     uint      `json:"amount"`
	Type       Type      `json:"type"`
	CategoryID entity.ID `json:"categoryID"`
	Category   Category  `json:"category"`
	UserID     entity.ID `json:"userID"`
}

func NewEntry(title string, amount uint, entryType Type, categoryID entity.ID, userID entity.ID) (*Entry, error) {
	if title == "" {
		return nil, custom_errors.ErrTitleRequired
	} else if len(title) < 2 {
		return nil, custom_errors.ErrTitleMinLength
	} else if len(title) > 100 {
		return nil, custom_errors.ErrTitleMaxLength
	}

	if amount == 0 {
		return nil, custom_errors.ErrAmountRequired
	}

	if entryType != In && entryType != Out {
		return nil, custom_errors.ErrInvalidType
	}

	if categoryID == uuid.Nil {
		return nil, custom_errors.ErrCategoryRequired
	}

	if userID == uuid.Nil {
		return nil, custom_errors.ErrUserIDRequired
	}

	return &Entry{
		Title:      title,
		Amount:     amount,
		Type:       entryType,
		CategoryID: categoryID,
		UserID:     userID,
	}, nil
}
