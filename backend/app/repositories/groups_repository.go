package repositories

import (
	"cupcake/app/database"
	domain "cupcake/app/models"
	"fmt"
)

type GroupsRepository interface {
	Insert(group *domain.Groups) (*domain.Groups, error)
	Find(id string) (*domain.Groups, error)
	FindAll() ([]*domain.Groups, error)
	Update(group *domain.Groups) (*domain.Groups, error)
}

type GroupsRepositoryDb struct {
	Db *database.Database
}

func (repo GroupsRepositoryDb) Insert(bet *domain.Groups) (*domain.Groups, error) {
	err := repo.Db.Create(bet).Error

	if err != nil {
		return nil, err
	}

	return bet, nil
}

func (repo GroupsRepositoryDb) Find(id string) (*domain.Groups, error) {
	var group domain.Groups

	repo.Db.First(&group, "id = ?", id)

	if group.UserID == "" {
		return nil, fmt.Errorf("group does not exist")
	}

	return &group, nil
}

func (repo GroupsRepositoryDb) FindAll() (*[]domain.Groups, error) {
	var bets []domain.Groups

	repo.Db.Find(&bets)

	return &bets, nil
}

func (repo GroupsRepositoryDb) Update(bet *domain.Groups) (*domain.Groups, error) {
	err := repo.Db.Save(&bet).Error

	if err != nil {
		return nil, err
	}

	return bet, nil
}

func (repo GroupsRepositoryDb) Delete(bet *domain.Groups) (*domain.Groups, error) {
	err := repo.Db.Delete(&bet).Error

	if err != nil {
		return nil, err
	}

	return bet, nil
}
