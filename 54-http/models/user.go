package models

import (
	"encoding/json"
	"errors"
)

type User struct {
	Id          uint   `json:"id" yaml:"id"`
	Name        string `json:"name" yaml:"name"`
	Email       string `json:"email"`
	Mobile      string `json:"mobile"`
	Status      string `json:"status"`
	LastModifed uint64 `json:"last_modified"`
}

func (u *User) Validate() error {
	if u.Name == "" {
		return errors.New("invalid name field")
	}
	if u.Email == "" {
		return errors.New("invalid email field")
	}
	if u.Mobile == "" {
		return errors.New("invalid mobile field")
	}
	return nil
}

func (u *User) ToSting() (string, error) {
	bytes, err := json.Marshal(u)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
