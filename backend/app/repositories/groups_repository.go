package repositories

import (
	"cupcake/app/database"
	"cupcake/app/models"
	"fmt"
)

type GroupsRepository interface {
	Insert(group *models.Groups) (*models.Groups, error)
	Find(id string) (*models.Groups, error)
	FindAll() ([]*models.Groups, error)
	Update(group *models.Groups) (*models.Groups, error)
}

type GroupsRepositoryDb struct {
	Db *database.Database
}

func (repo GroupsRepositoryDb) Insert(bet *models.Groups) (*models.Groups, error) {
	err := repo.Db.Create(bet).Error

	if err != nil {
		return nil, err
	}

	return bet, nil
}

func (repo GroupsRepositoryDb) Find(id string) (*models.Groups, error) {
	var group models.Groups

	repo.Db.First(&group, "id = ?", id)

	if group.UserID == "" {
		return nil, fmt.Errorf("group does not exist")
	}

	return &group, nil
}

func (repo GroupsRepositoryDb) FindAll() (*[]models.Groups, error) {
	var bets []models.Groups

	repo.Db.Find(&bets)

	return &bets, nil
}

func (repo GroupsRepositoryDb) Update(bet *models.Groups) (*models.Groups, error) {
	err := repo.Db.Save(&bet).Error

	if err != nil {
		return nil, err
	}

	return bet, nil
}

func (repo GroupsRepositoryDb) Delete(bet *models.Groups) (*models.Groups, error) {
	err := repo.Db.Delete(&bet).Error

	if err != nil {
		return nil, err
	}

	return bet, nil
}
