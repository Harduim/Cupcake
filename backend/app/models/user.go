package models

import (
	_ "github.com/go-playground/validator/v10"
	uuid "github.com/satori/go.uuid"
)

type User struct {
	ID      string `json:"id" validate:"required,uuid" gorm:"type:uuid;primary_key"`
	Name    string `json:"name" validate:"required" gorm:"type:varchar(255)"`
	Email   string `json:"email" validate:"required,email" gorm:"type:varchar(255)"`
	IsAdmin *bool  `json:"isAdmin" validate:"required" gorm:"type:bool;default:false"`
	Points  *int   `gorm:"default:0"`
}

func NewUser(name string, email string, isAdmin *bool) (*User, error) {
	user := &User{Name: name, Email: email, IsAdmin: isAdmin}
	user.prepare()

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

func (user *User) prepare() {
	user.ID = uuid.NewV4().String()
}
