package repositories

import (
	"cupcake/app/database"
	"cupcake/app/domain"
	"fmt"
)

type BracketRepository interface {
	Insert(bracket *domain.Bracket) (*domain.Bracket, error)
	Find(id string) (*domain.Bracket, error)
	FindAll() ([]*domain.Bracket, error)
	Update(bracket *domain.Bracket) (*domain.Bracket, error)
}

type BracketRepositoryDb struct {
	Db *database.Database
}

func (repo BracketRepositoryDb) Insert(bracket *domain.Bracket) (*domain.Bracket, error) {
	err := repo.Db.Create(bracket).Error

	if err != nil {
		return nil, err
	}

	return bracket, nil
}

func (repo BracketRepositoryDb) Find(id string) (*domain.Bracket, error) {
	var bracket domain.Bracket

	repo.Db.First(&bracket, "id = ?", id)

	if bracket.ID == "" {
		return nil, fmt.Errorf("bracket does not exist")
	}

	return &bracket, nil
}

func (repo BracketRepositoryDb) FindAll() (*[]domain.Bracket, error) {
	var brackets []domain.Bracket

	repo.Db.Find(&brackets)

	return &brackets, nil
}

func (repo BracketRepositoryDb) Update(bracket *domain.Bracket) (*domain.Bracket, error) {
	err := repo.Db.Save(&bracket).Error

	if err != nil {
		return nil, err
	}

	return bracket, nil
}
