package repositories

import (
	"cupcake/app/database"
	"cupcake/app/models"
	"fmt"
)

type BetRepository interface {
	Insert(match *models.Bet) (*models.Bet, error)
	Find(id string) (*models.Bet, error)
	FindAll() ([]*models.Bet, error)
	Update(match *models.Bet) (*models.Bet, error)
}

type BetRepositoryDb struct {
	Db *database.Database
}

func (repo BetRepositoryDb) Insert(bet *models.Bet) (*models.Bet, error) {
	err := repo.Db.Create(bet).Error

	if err != nil {
		return nil, err
	}

	return bet, nil
}

func (repo BetRepositoryDb) Find(id string) (*models.Bet, error) {
	var match models.Bet

	repo.Db.First(&match, "id = ?", id)

	if match.ID == "" {
		return nil, fmt.Errorf("match does not exist")
	}

	return &match, nil
}

func (repo BetRepositoryDb) FindAll() (*[]models.Bet, error) {
	var bets []models.Bet

	repo.Db.Find(&bets)

	return &bets, nil
}

func (repo BetRepositoryDb) Update(bet *models.Bet) (*models.Bet, error) {
	err := repo.Db.Save(&bet).Error

	if err != nil {
		return nil, err
	}

	return bet, nil
}

func (repo BetRepositoryDb) Delete(bet *models.Bet) (*models.Bet, error) {
	err := repo.Db.Delete(&bet).Error

	if err != nil {
		return nil, err
	}

	return bet, nil
}
