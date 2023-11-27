package entity

import (
	"strings"
	"testing"

	"github.com/oaraujocesar/mpf/pkg/custom_errors"
	"github.com/stretchr/testify/assert"
)

func assertUserCreation(t *testing.T, name, avatar, email, password string, role Role) *User {
	user, err := NewUser(name, avatar, email, password, role)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.Equal(t, name, user.Name)
	assert.Equal(t, avatar, user.Avatar)
	assert.Equal(t, email, user.Email)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, role, user.Role)
	return user
}

func TestNewUserAdmin(t *testing.T) {
	assertUserCreation(t, "Cesar", "https://www.gravatar.com/avatar/205e460b479e2e5b48aec07710c08d50", "cesar@admin.com", "reallystrongpassword", Admin)
}

func TestNewUserCommon(t *testing.T) {
	assertUserCreation(t, "Cesar", "https://www.gravatar.com/avatar/205e460b479e2e5b48aec07710c08d50", "cesar@admin.com", "reallystrongpassword", Common)
}

func TestNewUserWithInvalidRole(t *testing.T) {
	user, err := NewUser("Cesar", "https://www.gravatar.com/avatar/205e460b479e2e5b48aec07710c08d50", "cesar@admin.com", "reallystrongpassword", "")
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.Equal(t, custom_errors.ErrInvalidRole, err)
}

func TestNewUserWithEmptyName(t *testing.T) {
	user, err := NewUser("", "https://www.gravatar.com/avatar/205e460b479e2e5b48aec07710c08d50", "cesar@user.com", "reallystrongpassword", Common)
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.Equal(t, custom_errors.ErrNameRequired, err)
}

func TestNewUserWithInvalidNameLength(t *testing.T) {
	user, err := NewUser("C", "https://www.gravatar.com/avatar/205e460b479e2e5b48aec07710c08d50", "cesar@user", "reallystrongpassword", Common)
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.Equal(t, custom_errors.ErrNameMinLength, err)
}

func TestNewUserWithInvalidNameMaxLength(t *testing.T) {
	user, err := NewUser(strings.Repeat("a", 101), "https://www.gravatar.com/avatar/205e460b479e2e5b48aec07710c08d50", "cesar@user", "reallystrongpassword", Common)

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.Equal(t, custom_errors.ErrNameMaxLength, err)
}

func TestNewUserWithEmptyEmail(t *testing.T) {
	user, err := NewUser("Cesar", "https://www.gravatar.com/avatar/205e460b479e2e5b48aec07710c08d50", "", "reallystrongpassword", Common)
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.Equal(t, custom_errors.ErrEmailRequired, err)
}

func TestNewUserWithEmptyPassword(t *testing.T) {
	user, err := NewUser("Cesar", "https://www.gravatar.com/avatar/205e460b479e2e5b48aec07710c08d50", "cesar@user", "", Common)
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.Equal(t, custom_errors.ErrPasswordRequired, err)
}

func TestNewUserWithWrongPassword(t *testing.T) {
	user, err := NewUser("Cesar", "https://www.gravatar.com/avatar/205e460b479e2e5b48aec07710c08d50", "cesar@user", "weak", Common)
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.Equal(t, custom_errors.ErrPasswordMinLength, err)
}

func TestUserValidatePassword(t *testing.T) {
	user := assertUserCreation(t, "Cesar", "https://www.gravatar.com/avatar/205e460b479e2e5b48aec07710c08d50", "cesar@admin.com", "reallystrongpassword", Admin)
	assert.True(t, user.ValidatePassword("reallystrongpassword"))
	assert.False(t, user.ValidatePassword("reallyweakpassword"))
	assert.NotEqual(t, "reallystrongpassword", user.Password)
}
