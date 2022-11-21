package models

import (
	_ "github.com/go-playground/validator/v10"
)

type User struct {
	ID      string `json:"id" gorm:"type:uuid;primary_key"`
	Name    string `json:"name" validate:"required" gorm:"type:varchar(255)"`
	IsAdmin *bool  `json:"isAdmin" validate:"required" gorm:"type:bool;default:false"`
	Points  *int   `json:"points" gorm:"default:0"`
}

func NewUser(name string, oid string, isAdmin *bool) (*User, error) {
	user := &User{ID: oid, Name: name, IsAdmin: isAdmin}

	err := user.Validate()

	if err != nil {
		return nil, err
	}

	return user, nil

}

func (user *User) Validate() error {
	err := validate.Struct(user)

	if err != nil {
		return err
	}
	return nil
}
