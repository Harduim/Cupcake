package repositories

import (
	"cupcake/app/database"
	"cupcake/app/models"
	"fmt"
)

type GroupsRepository interface {
	Insert(group *models.Group) (*models.Group, error)
	Find(id string) (*models.Group, error)
	FindAll() ([]*models.Group, error)
	Update(group *models.Group) (*models.Group, error)
}

type GroupsRepositoryDb struct {
	Db *database.Database
}

func (repo GroupsRepositoryDb) Insert(bet *models.Group) (*models.Group, error) {
	err := repo.Db.Create(bet).Error

	if err != nil {
		return nil, err
	}

	return bet, nil
}

func (repo GroupsRepositoryDb) Find(id string) (*models.Group, error) {
	var group models.Group

	repo.Db.First(&group, "id = ?", id)

	if group.UserID == "" {
		return nil, fmt.Errorf("group does not exist")
	}

	return &group, nil
}

func (repo GroupsRepositoryDb) FindAll() (*[]models.Group, error) {
	var groups []models.Group

	repo.Db.Find(&groups)

	return &groups, nil
}

func (repo GroupsRepositoryDb) Update(bet *models.Group) (*models.Group, error) {
	err := repo.Db.Save(&bet).Error

	if err != nil {
		return nil, err
	}

	return bet, nil
}

func (repo GroupsRepositoryDb) Delete(bet *models.Group) (*models.Group, error) {
	err := repo.Db.Delete(&bet).Error

	if err != nil {
		return nil, err
	}

	return bet, nil
}
