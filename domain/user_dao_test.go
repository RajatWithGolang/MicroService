package domain

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNoUserFound(t *testing.T) {
	user, err := UserDAO.GetUser(0)
	assert.Nil(t, user, "Not expecting the user with ID 0")
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusNotFound, err.StatusCode)
	assert.Equal(t, "Not found", err.Code)
	assert.Equal(t, "user with Id 0 not found", err.Message)

}
func TestGetUserNoError(t *testing.T) {
	user, err := UserDAO.GetUser(123)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, 123, user.Id)
	assert.Equal(t, "Rajat", user.Name)
	assert.Equal(t, "rajat@test.com", user.Email)
}
