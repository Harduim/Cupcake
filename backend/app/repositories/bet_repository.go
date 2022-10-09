package repositories

import (
	"cupcake/app/database"
	"cupcake/app/domain"
	"fmt"
)

type BetRepository interface {
	Insert(match *domain.Bet) (*domain.Bet, error)
	Find(id string) (*domain.Bet, error)
	Update(match *domain.Bet) (*domain.Bet, error)
}

type BetRepositoryDb struct {
	Db *database.Database
}

func (bet BetRepositoryDb) Insert(match *domain.Bet) (*domain.Bet, error) {
	err := bet.Db.Create(match).Error

	if err != nil {
		return nil, err
	}

	return match, nil
}

func (bet BetRepositoryDb) Find(id string) (*domain.Bet, error) {
	var match domain.Bet

	bet.Db.First(&match, "id = ?", id)

	if match.ID == "" {
		return nil, fmt.Errorf("match does not exist")
	}

	return &match, nil
}

func (bet BetRepositoryDb) Update(match *domain.Bet) (*domain.Bet, error) {
	err := bet.Db.Save(&match).Error

	if err != nil {
		return nil, err
	}

	return match, nil
}
