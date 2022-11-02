package repositories

import (
	"cupcake/app/database"
	"cupcake/app/domain"
	"fmt"
)

type JokerRepository interface {
	Insert(joker *domain.Joker) (*domain.Joker, error)
	Find(id string) (*domain.Joker, error)
	FindAll() ([]*domain.Joker, error)
	Update(joker *domain.Joker) (*domain.Joker, error)
}

type JokerRepositoryDb struct {
	Db *database.Database
}

func (repo JokerRepositoryDb) Insert(joker *domain.Joker) (*domain.Joker, error) {
	err := repo.Db.Create(joker).Error

	if err != nil {
		return nil, err
	}

	return joker, nil
}

func (repo JokerRepositoryDb) Find(id string) (*domain.Joker, error) {
	var joker domain.Joker

	repo.Db.First(&joker, "id = ?", id)

	if joker.ID == "" {
		return nil, fmt.Errorf("joker does not exist")
	}

	return &joker, nil
}

func (repo JokerRepositoryDb) FindAll() (*[]domain.Joker, error) {
	var bets []domain.Joker

	repo.Db.Find(&bets)

	return &bets, nil
}

func (repo JokerRepositoryDb) Update(bet *domain.Joker) (*domain.Joker, error) {
	err := repo.Db.Save(&bet).Error

	if err != nil {
		return nil, err
	}

	return bet, nil
}

func (repo JokerRepositoryDb) Delete(bet *domain.Joker) (*domain.Joker, error) {
	err := repo.Db.Delete(&bet).Error

	if err != nil {
		return nil, err
	}

	return bet, nil
}
