package domain

type UserPoints struct {
	User   *User  `json:"user" valid:"-"`
	UserID string `json:"-" validate:"-" gorm:"type:uuid;notnull"`
	Points int8   `json:"points"`
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
