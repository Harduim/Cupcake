package models

type Key struct {
	ID   string `json:"id" validate:"required,uuid" gorm:"type:uuid;primary_key"`
	Name string `json:"name" validate:"required" gorm:"type:varchar(255)"`
}

func NewKey() *Key {
	return &Key{}
}

func (key *Key) Validate() error {
	err := validate.Struct(key)

	if err != nil {
		return err
	}
	return nil
}
