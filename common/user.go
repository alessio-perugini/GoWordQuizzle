package common

import (
	"errors"
	"github.com/tevino/abool"
	"sync"
)

type User struct {
	Score      int64
	nickname   string
	password   string
	friendList map[string]string
	UdpPort    uint16
	Connected  bool
	InGame     abool.AtomicBool
}

func NewUser() *User {
	return &User{
		UdpPort: 0,
	}
}

func (u *User) GetNickname() string {
	return u.nickname
}

func (u *User) SetNickname(name string) {
	u.nickname = name
}

func (u *User) GetPassword() string {
	return u.password
}

func (u *User) SetPssword(pw string) {
	u.password = pw
}

type UsersList struct {
	users map[string]User //TODO handle concurrency with mutex or sync stuff
}

var instanceUserList *UsersList
var once sync.Once

func GetInstanceUsersList() *UsersList {
	//TODO load data from json
	once.Do(func() {
		ls := make(map[string]User, 0)
		instanceUserList = &UsersList{users: ls}
	})

	return instanceUserList
}

func (ul *UsersList) GetUsers() map[string]User {
	return ul.users
}

func (ul *UsersList) AddUser(u User) error {
	if u.nickname == "" {
		return errors.New("User not valid")
	}
	if ul.users[u.nickname].nickname == u.nickname {
		return errors.New("User already exists")
	}

	ul.users[u.nickname] = u
	return nil
}

func (ul *UsersList) RemoveUser(u User) error {
	if u.nickname == "" {
		return errors.New("User not valid")
	}
	if ul.users[u.nickname].nickname != u.nickname {
		return errors.New("User not exists")
	}

	delete(ul.users, u.nickname)

	return nil
}

func (ul *UsersList) IsConnected(nickname string) (bool, error) {
	if nickname == "" {
		return false, errors.New("User not valid")
	}
	if ul.users[nickname].nickname != nickname {
		return false, errors.New("User not exists")
	}

	return ul.users[nickname].Connected, nil
}

func (ul *UsersList) SetConnected(nickname string, value bool) error {
	if nickname == "" {
		return errors.New("User not valid")
	}
	if ul.users[nickname].nickname != nickname {
		return errors.New("User not exists")
	}

	user := ul.users[nickname]
	user.Connected = value
	ul.users[nickname] = user
	return nil
}

func (ul *UsersList) GetUser(nickname string) (User, error) {
	if nickname == "" {
		return User{}, errors.New("User not valid")
	}

	user := ul.users[nickname]
	if user.nickname != nickname {
		return User{}, errors.New("User not exists")
	}

	return user, nil
}
