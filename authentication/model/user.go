package model

import (
	"strings"
)

type (
	User struct {
		ID       int
		Name     string
		Email    string
		Salt     string
		Password string
	}
)

func (u *User) GetFirstName() string {
	splitName := strings.Split(u.Name, " ")
	return strings.Join(splitName[:len(splitName)-1], " ")
}

func (u *User) GetLastName() string {
	splitName := strings.Split(u.Name, " ")

	return splitName[len(splitName)-1]
}
