package model

import "testing"

func TestUser(t *testing.T) *User {
	return &User{
		Email:    "durov@mail.ru",
		Password: "telega_top",
	}
}
