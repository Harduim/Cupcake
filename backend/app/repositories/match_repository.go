package repositories

import (
	"cupcake/app/database"
	"cupcake/app/domain"
	"fmt"
)

type MatchRepository interface {
	Insert(match *domain.Match) (*domain.Match, error)
	Find(id string) (*domain.Match, error)
	FindAll() ([]*domain.Match, error)
	Update(match *domain.Match) (*domain.Match, error)
}

type MatchRepositoryDb struct {
	Db *database.Database
}

func (repo MatchRepositoryDb) Insert(match *domain.Match) (*domain.Match, error) {
	err := repo.Db.Create(match).Error

	if err != nil {
		return nil, err
	}

	return match, nil
}

func (repo MatchRepositoryDb) Find(id string) (*domain.Match, error) {
	var match domain.Match

	repo.Db.First(&match, "id = ?", id)

	if match.ID == "" {
		return nil, fmt.Errorf("match does not exist")
	}

	return &match, nil
}

func (repo MatchRepositoryDb) FindAll() (*[]domain.Match, error) {
	var matches []domain.Match

	repo.Db.Find(&matches)

	return &matches, nil
}

func (repo MatchRepositoryDb) Update(match *domain.Match) (*domain.Match, error) {
	err := repo.Db.Save(&match).Error

	if err != nil {
		return nil, err
	}

	return match, nil
}
