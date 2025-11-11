package week2

import (
	"errors"
	"strings"
)

type User struct {
	Name  string
	Email string
	Age   int
}

func (u User) Validate() error {
	if u.Name == "" {
		return errors.New("имя не должно быть пустым")
	}
	if u.Age < 0 || u.Age > 150 {
		return errors.New("возраст должен быть от 0 до 150")
	}
	if !strings.Contains(u.Email, "@") {
		return errors.New("email должен содержать @")
	}
	return nil
}
