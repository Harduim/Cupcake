package domain

type UserPoints struct {
	User   *User  `gorm:"foreignKey:UserID"`
	UserID string `json:"user_id" validate:"required,uuid" gorm:"type:varchar(255);primary_key"`
	Points int    `json:"points" validate:"required" gorm:"type:integer"`
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
