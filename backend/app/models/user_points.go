package models

type UserPoints struct {
	User   User `gorm:"embedded"`
	Points int8 `json:"points"`
}

func NewUserPoints() *UserPoints {
	return &UserPoints{}
}

func (userPoints *UserPoints) Validate() error {
	err := validate.Struct(userPoints)

	if err != nil {
		return err
	}
	return nil
}
