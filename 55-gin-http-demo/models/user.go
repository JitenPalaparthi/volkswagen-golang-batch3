package models

import (
	"encoding/json"
	"errors"
)

func init() {
	println("calling init in models-1")
}

func init() {
	println("calling init in models-2")
}

type User struct {
	//gorm.Model
	Id          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Email       string `json:"email" gorm:"unique"`
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
