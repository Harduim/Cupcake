package repository

import (
	"cupcake/app/domain"
	"fmt"
	"gorm.io/gorm"
)

type UserRepository interface {
	Insert(user *domain.User) (*domain.User, error)
	Find(id string) (*domain.User, error)
	Update(user *domain.User) (*domain.User, error)
}

type UserRepositoryDb struct {
	Db *gorm.DB
}

func (repo UserRepositoryDb) Insert(user *domain.User) (*domain.User, error) {
	err := repo.Db.Create(user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repo UserRepositoryDb) Find(id string) (*domain.User, error) {
	var user domain.User

	repo.Db.First(&user, "id = ?", id)

	if user.ID == "" {
		return nil, fmt.Errorf("job does not exist")
	}

	return &user, nil
}

func (repo UserRepositoryDb) Update(user *domain.User) (*domain.User, error) {
	err := repo.Db.Save(&user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}
