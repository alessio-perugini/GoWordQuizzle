package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetIstanceUsersList(t *testing.T){
	i1 := GetInstanceUsersList()
	i2 := GetInstanceUsersList()

	assert.Equal(t, i1, i2)
}

func TestUsersList_AddUser(t *testing.T) {
	u1 := User{nickname: "gino"}

	assert.NoError(t, GetInstanceUsersList().AddUser(u1))
	assert.EqualError(t, GetInstanceUsersList().AddUser(u1), "User already exists")
	assert.EqualError(t, GetInstanceUsersList().AddUser(User{nickname: ""}), "User not valid")
}

func TestUsersList_RemoveUser(t *testing.T) {
	u1 := User{nickname: "gino"}

	assert.EqualError(t, GetInstanceUsersList().RemoveUser(User{nickname: ""}), "User not valid")
	assert.EqualError(t, GetInstanceUsersList().RemoveUser(u1), "User not exists")

	GetInstanceUsersList().AddUser(u1)
	assert.Nil(t, GetInstanceUsersList().RemoveUser(u1))
}

func TestUsersList_IsConnected(t *testing.T) {
	nickname := "gino"
	u1 := User{nickname: nickname, Connected: true}
	GetInstanceUsersList().AddUser(u1)

	if connected, err := GetInstanceUsersList().IsConnected(nickname); err == nil {
		assert.True(t, connected)
	}

	GetInstanceUsersList().SetConnected(nickname, false)
	if connected, err := GetInstanceUsersList().IsConnected(nickname); err == nil {
		assert.False(t, connected)
	}

	if _, err := GetInstanceUsersList().IsConnected(""); err != nil {
		assert.EqualError(t, err, "User not valid")
	}

	if _, err := GetInstanceUsersList().IsConnected("pina"); err != nil {
		assert.EqualError(t, err, "User not exists")
	}
}

func TestUsersList_GetUser(t *testing.T) {
	nickname := "gino"

	if _, err := GetInstanceUsersList().GetUser(""); err != nil {
		assert.EqualError(t, err, "User not valid")
	}
	if _, err := GetInstanceUsersList().GetUser(nickname); err != nil {
		assert.EqualError(t, err, "User not exists")
	}

	u1 := User{nickname: nickname}
	GetInstanceUsersList().AddUser(u1)

	if value, err := GetInstanceUsersList().GetUser(nickname); err != nil {
		assert.Equal(t, value, u1)
	}
}

func TestUsersList_SetConnected(t *testing.T) {
	nickname := "gino"
	u1 := User{nickname: nickname, Connected: true}

	assert.EqualError(t, GetInstanceUsersList().SetConnected( "", false), "User not valid")
	assert.EqualError(t, GetInstanceUsersList().SetConnected( nickname, false), "User not exists")

	GetInstanceUsersList().AddUser(u1)
	assert.Nil(t, GetInstanceUsersList().SetConnected(nickname, false))
	u2, _ := GetInstanceUsersList().GetUser(nickname)
	assert.False(t, u2.Connected)

}

