package entity

import (
	"strings"
	"testing"

	"github.com/oaraujocesar/mpf/pkg/custom_errors"
	"github.com/stretchr/testify/assert"
)

func TestNewCategory(t *testing.T) {
	category, err := NewCategory("Pets")

	assert.Nil(t, err)
	assert.NotNil(t, category)
	assert.NotEmpty(t, category.ID)
	assert.Equal(t, "Pets", category.Name)
}

func TestNewCategoryWithEmptyName(t *testing.T) {
	category, err := NewCategory("")

	assert.Nil(t, category)
	assert.NotNil(t, err)
	assert.Equal(t, custom_errors.ErrNameRequired, err)
}

func TestNewCategoryWithInvalidNameLength(t *testing.T) {
	category, err := NewCategory("P")

	assert.Nil(t, category)
	assert.NotNil(t, err)
	assert.Equal(t, custom_errors.ErrNameMinLength, err)
}

func TestNewCategoryWithInvalidNameMaxLength(t *testing.T) {
	category, err := NewCategory(strings.Repeat("a", 101))

	assert.Nil(t, category)
	assert.NotNil(t, err)
	assert.Equal(t, custom_errors.ErrNameMaxLength, err)
}
