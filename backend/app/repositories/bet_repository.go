package repositories

import (
	"cupcake/app/database"
	"cupcake/app/models"
	"fmt"
)

type BetRepository interface {
	Insert(bet *models.Bet) (*models.Bet, error)
	Find(id string) (*models.Bet, error)
	FindAll() ([]*models.Bet, error)
	Update(bet *models.Bet) (*models.Bet, error)
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
	var bet models.Bet

	repo.Db.First(&bet, "id = ?", id)

	if bet.ID == "" {
		return nil, fmt.Errorf("match does not exist")
	}

	return &bet, nil
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
