package server

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLogin(t *testing.T) {
	username := "gino"
	pw := "superstrongpw"

	assert.EqualError(t, login("", pw), "username not valid")
	assert.EqualError(t, login(username, ""), "password not valid")

	login(username, pw)
	//Test user already logged
	assert.EqualError(t, login(username, pw), "User already exists")
}
