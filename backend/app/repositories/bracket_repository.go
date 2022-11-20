package repositories

import (
	"cupcake/app/database"
	"cupcake/app/models"
	"fmt"

	"gorm.io/gorm/clause"
)

type BracketRepository interface {
	Insert(bracket *models.Bracket) (*models.Bracket, error)
	Find(id string) (*models.Bracket, error)
	FindAll() ([]*models.Bracket, error)
	Update(bracket *models.Bracket) (*models.Bracket, error)
}

type BracketRepositoryDb struct {
	Db *database.Database
}

func (repo BracketRepositoryDb) Insert(bracket *models.Bracket) (*models.Bracket, error) {
	err := repo.Db.Create(bracket).Error

	if err != nil {
		return nil, err
	}

	return bracket, nil
}

func (repo BracketRepositoryDb) Find(id string) (*models.Bracket, error) {
	var bracket models.Bracket

	repo.Db.First(&bracket, "id = ?", id)

	if bracket.ID == "" {
		return nil, fmt.Errorf("bracket does not exist")
	}

	return &bracket, nil
}

func (repo BracketRepositoryDb) FindAll() (*[]models.Bracket, error) {
	var brackets []models.Bracket

	repo.Db.Preload("Matches.NationalTeamA").Preload("Matches.NationalTeamB").Preload("Matches.Winner").Preload(clause.Associations).Find(&brackets)

	return &brackets, nil
}

func (repo BracketRepositoryDb) Update(bracket *models.Bracket) (*models.Bracket, error) {
	err := repo.Db.Save(&bracket).Error

	if err != nil {
		return nil, err
	}

	return bracket, nil
}
