package repositories

import (
	"cupcake/app/database"
	"cupcake/app/domain"
	"fmt"
)

type UserPointsRepository interface {
	Insert(userPoints *domain.UserPoints) (*domain.UserPoints, error)
	Find(id string) (*domain.UserPoints, error)
	FindAll() ([]*domain.UserPoints, error)
	Update(userPoints *domain.UserPoints) (*domain.UserPoints, error)
}

type UserPointsRepositoryDb struct {
	Db *database.Database
}

func (repo UserPointsRepositoryDb) Insert(userPoints *domain.UserPoints) (*domain.UserPoints, error) {
	err := repo.Db.Create(userPoints).Error

	if err != nil {
		return nil, err
	}

	return userPoints, nil
}

func (repo UserPointsRepositoryDb) Find(id string) (*domain.UserPoints, error) {
	var userPoints domain.UserPoints

	repo.Db.First(&userPoints, "user_id = ?", id)

	if userPoints.UserID == "" {
		return nil, fmt.Errorf("user does not have points")
	}

	return &userPoints, nil
}

func (repo UserPointsRepositoryDb) FindAll() (*[]domain.UserPoints, error) {
	var userPoints []domain.UserPoints

	repo.Db.Find(&userPoints)

	return &userPoints, nil
}

func (repo UserPointsRepositoryDb) Update(userPoints *domain.UserPoints) (*domain.UserPoints, error) {
	err := repo.Db.Save(&userPoints).Error

	if err != nil {
		return nil, err
	}

	return userPoints, nil
}
