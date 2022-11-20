package repositories

import (
	"cupcake/app/database"
	"cupcake/app/models"
	"fmt"
)

type JokerRepository interface {
	Insert(joker *models.Joker) (*models.Joker, error)
	Find(id string) (*models.Joker, error)
	FindAll() ([]*models.Joker, error)
	Update(joker *models.Joker) (*models.Joker, error)
}

type JokerRepositoryDb struct {
	Db *database.Database
}

func (repo JokerRepositoryDb) Insert(joker *models.Joker) (*models.Joker, error) {
	err := repo.Db.Create(joker).Error

	if err != nil {
		return nil, err
	}

	return joker, nil
}

func (repo JokerRepositoryDb) Find(id string) (*models.Joker, error) {
	var joker models.Joker

	repo.Db.First(&joker, "id = ?", id)

	if joker.ID == "" {
		return nil, fmt.Errorf("joker does not exist")
	}

	return &joker, nil
}

func (repo JokerRepositoryDb) FindAll() (*[]models.Joker, error) {
	var bets []models.Joker

	repo.Db.Find(&bets)

	return &bets, nil
}

func (repo JokerRepositoryDb) Update(bet *models.Joker) (*models.Joker, error) {
	err := repo.Db.Save(&bet).Error

	if err != nil {
		return nil, err
	}

	return bet, nil
}

func (repo JokerRepositoryDb) Delete(bet *models.Joker) (*models.Joker, error) {
	err := repo.Db.Delete(&bet).Error

	if err != nil {
		return nil, err
	}

	return bet, nil
}
