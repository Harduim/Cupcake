package domain

import (
	_ "github.com/go-playground/validator/v10"
)

type User struct {
	ID      string `json:"id" validate:"required,uuid" gorm:"type:uuid;primary_key"`
	Name    string `json:"name" validate:"required" gorm:"type:varchar(255)"`
	Email   string `json:"email" validate:"required,email" gorm:"type:varchar(255)"`
	IsAdmin *bool  `json:"is_admin" validate:"required" gorm:"type:bool;default:false"`
}

func NewUser() *User {
	return &User{}
}

func (user *User) Validate() error {
	err := validate.Struct(user)

	if err != nil {
		return err
	}
	return nil
}
