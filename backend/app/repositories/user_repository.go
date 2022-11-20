package repositories

import (
	"cupcake/app/database"
	"cupcake/app/models"
	"fmt"
)

type UserRepository interface {
	Insert(user *models.User) (*models.User, error)
	Find(id string) (*models.User, error)
	FindAll() ([]*models.User, error)
	Update(user *models.User) (*models.User, error)
}

type UserRepositoryDb struct {
	Db *database.Database
}

func (repo UserRepositoryDb) Insert(user *models.User) (*models.User, error) {
	err := repo.Db.Create(user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repo UserRepositoryDb) Find(id string) (*models.User, error) {
	var user models.User

	repo.Db.First(&user, "id = ?", id)

	if user.ID == "" {
		return nil, fmt.Errorf("user does not exist")
	}

	return &user, nil
}

func (repo UserRepositoryDb) Update(user *models.User) (*models.User, error) {
	err := repo.Db.Save(&user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repo UserRepositoryDb) FindAll() (*[]models.User, error) {
	var users []models.User

	repo.Db.Find(&users)

	return &users, nil
}

func (repo UserRepositoryDb) FindByEmail(id string) (*models.User, error) {
	var user models.User

	repo.Db.First(&user, "email = ?", id)

	if user.Email == "" {
		return nil, fmt.Errorf("user does not exist")
	}

	return &user, nil
}
