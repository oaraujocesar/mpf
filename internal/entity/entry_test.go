package entity

import (
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/oaraujocesar/mpf/pkg/custom_errors"
	"github.com/oaraujocesar/mpf/pkg/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewEntryValidIn(t *testing.T) {
	title := "Valid Entry"
	amount := uint(100)
	entryType := In
	categoryID := entity.NewID()
	userID := entity.NewID()

	entry, err := NewEntry(title, amount, entryType, categoryID, userID)

	assert.Nil(t, err)
	assert.NotNil(t, entry)
	assert.Equal(t, title, entry.Title)
	assert.Equal(t, amount, entry.Amount)
	assert.Equal(t, entryType, entry.Type)
	assert.Equal(t, categoryID, entry.CategoryID)
	assert.Equal(t, userID, entry.UserID)
}

func TestNewEntryValidOut(t *testing.T) {
	title := "Valid Entry"
	amount := uint(100)
	entryType := Out
	categoryID := entity.NewID()
	userID := entity.NewID()

	entry, err := NewEntry(title, amount, entryType, categoryID, userID)

	assert.Nil(t, err)
	assert.NotNil(t, entry)
	assert.Equal(t, title, entry.Title)
	assert.Equal(t, amount, entry.Amount)
	assert.Equal(t, entryType, entry.Type)
	assert.Equal(t, categoryID, entry.CategoryID)
	assert.Equal(t, userID, entry.UserID)
}

func TestNewEntryInvalid(t *testing.T) {
	testCases := []struct {
		title         string
		amount        uint
		entryType     Type
		categoryID    uuid.UUID
		userID        uuid.UUID
		expectedError error
	}{
		{"", 100, In, entity.NewID(), entity.NewID(), custom_errors.ErrTitleRequired},
		{"L", 100, In, entity.NewID(), entity.NewID(), custom_errors.ErrTitleMinLength},
		{strings.Repeat("a", 101), 100, Out, entity.NewID(), entity.NewID(), custom_errors.ErrTitleMaxLength},
		{"test", 0, Out, entity.NewID(), entity.NewID(), custom_errors.ErrAmountRequired},
		{"test", 50, Type("test"), entity.NewID(), entity.NewID(), custom_errors.ErrInvalidType},
		{"test", 50, In, uuid.Nil, entity.NewID(), custom_errors.ErrCategoryRequired},
		{"test", 50, In, entity.NewID(), uuid.Nil, custom_errors.ErrUserIDRequired},
	}

	for _, tc := range testCases {
		entry, err := NewEntry(tc.title, tc.amount, tc.entryType, tc.categoryID, tc.userID)

		assert.Nil(t, entry)
		assert.Equal(t, tc.expectedError, err)
	}
}
