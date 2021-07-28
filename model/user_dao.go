package model

import (
	"errors"
	"fmt"
)

var (
	users = map[int64]*User{
		123: {Id: 123, FirstName: "Scott", LastName: "Fitzgerald", Email: "myemailaddress@gmail.com"},
	}
)

func GetUser(userId int64) (*User, error) {
	if user := users[userId]; user != nil {
		return user, nil
	}

	return nil, errors.New(fmt.Sprintf("user %v was not found", userId))
}
