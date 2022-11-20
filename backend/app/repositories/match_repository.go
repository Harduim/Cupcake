package repositories

import (
	"cupcake/app/database"
	"cupcake/app/models"
	"fmt"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm/clause"
)

type MatchRepository interface {
	Insert(match *models.Match) (*models.Match, error)
	Find(id string) (*models.Match, error)
	FindAll() ([]*models.Match, error)
	Update(match *models.Match) (*models.Match, error)
}

type MatchRepositoryDb struct {
	Db *database.Database
}

func (repo MatchRepositoryDb) Insert(match *models.Match) (*models.Match, error) {
	if match.ID == "" {
		match.ID = uuid.NewV4().String()
	}

	err := repo.Db.Create(match).Error

	if err != nil {
		return nil, err
	}

	return match, nil
}

func (repo MatchRepositoryDb) Find(id string) (*models.Match, error) {
	var match models.Match

	repo.Db.First(&match, "id = ?", id)

	if match.ID == "" {
		return nil, fmt.Errorf("match does not exist")
	}

	return &match, nil
}

func (repo MatchRepositoryDb) FindAll() (*[]models.Match, error) {
	var matches []models.Match

	err := repo.Db.Model(&models.Match{}).Preload(clause.Associations).Find(&matches).Error

	if err != nil {
		return nil, err
	}

	return &matches, nil
}

func (repo MatchRepositoryDb) Update(match *models.Match) (*models.Match, error) {
	err := repo.Db.Save(&match).Error

	if err != nil {
		return nil, err
	}

	return match, nil
}

func (repo MatchRepositoryDb) Delete(match *models.Match) (*models.Match, error) {
	err := repo.Db.Delete(&match).Error

	if err != nil {
		return nil, err
	}

	return match, nil
}
