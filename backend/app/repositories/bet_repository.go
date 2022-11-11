package repositories

import (
	"cupcake/app/database"
	"cupcake/app/domain"
	"fmt"
)

type BetRepository interface {
	Insert(match *domain.Bet) (*domain.Bet, error)
	Find(id string) (*domain.Bet, error)
	FindAll() ([]*domain.Bet, error)
	Update(match *domain.Bet) (*domain.Bet, error)
}

type BetRepositoryDb struct {
	Db *database.Database
}

func (repo BetRepositoryDb) Insert(bet *domain.Bet) (*domain.Bet, error) {
	err := repo.Db.Create(bet).Error

	if err != nil {
		return nil, err
	}

	return bet, nil
}

func (repo BetRepositoryDb) Find(id string) (*domain.Bet, error) {
	var match domain.Bet

	repo.Db.First(&match, "id = ?", id)

	if match.ID == "" {
		return nil, fmt.Errorf("match does not exist")
	}

	return &match, nil
}

func (repo BetRepositoryDb) FindAll() (*[]domain.Bet, error) {
	var bets []domain.Bet

	repo.Db.Find(&bets)

	return &bets, nil
}

func (repo BetRepositoryDb) Update(bet *domain.Bet) (*domain.Bet, error) {
	err := repo.Db.Save(&bet).Error

	if err != nil {
		return nil, err
	}

	return bet, nil
}

func (repo BetRepositoryDb) Delete(bet *domain.Bet) (*domain.Bet, error) {
	err := repo.Db.Delete(&bet).Error

	if err != nil {
		return nil, err
	}

	return bet, nil
}
