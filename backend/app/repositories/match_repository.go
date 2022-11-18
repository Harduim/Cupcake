package repositories

import (
	"cupcake/app/database"
	domain "cupcake/app/models"
	"fmt"

	uuid "github.com/satori/go.uuid"
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
	if match.ID == "" {
		match.ID = uuid.NewV4().String()
	}

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

	err := repo.Db.Model(&domain.Match{}).Preload("Bracket").Find(&matches).Error

	if err != nil {
		return nil, err
	}

	return &matches, nil
}

func (repo MatchRepositoryDb) Update(match *domain.Match) (*domain.Match, error) {
	err := repo.Db.Save(&match).Error

	if err != nil {
		return nil, err
	}

	return match, nil
}

func (repo MatchRepositoryDb) Delete(match *domain.Match) (*domain.Match, error) {
	err := repo.Db.Delete(&match).Error

	if err != nil {
		return nil, err
	}

	return match, nil
}
