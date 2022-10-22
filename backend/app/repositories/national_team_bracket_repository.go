package repositories

import (
	"cupcake/app/database"
	"cupcake/app/domain"
	"fmt"
)

type NationalTeamBracketRepository interface {
	Insert(nationalTeamBracket *domain.NationalTeamBracket) (*domain.NationalTeamBracket, error)
	Find(id string) (*domain.NationalTeamBracket, error)
	Update(nationalTeamBracket *domain.NationalTeamBracket) (*domain.NationalTeamBracket, error)
}

type NationalTeamBracketRepositoryDb struct {
	Db *database.Database
}

func (repo NationalTeamBracketRepositoryDb) Insert(nationalTeamBracket *domain.NationalTeamBracket) (*domain.NationalTeamBracket, error) {
	err := repo.Db.Create(nationalTeamBracket).Error

	if err != nil {
		return nil, err
	}

	return nationalTeamBracket, nil
}

func (repo NationalTeamBracketRepositoryDb) Find(id string) (*domain.NationalTeamBracket, error) {
	var nationalTeamBracket domain.NationalTeamBracket

	repo.Db.First(&nationalTeamBracket, "national_team_id = ?", id)

	if nationalTeamBracket.BracketID == "" || nationalTeamBracket.NationalTeamID == "" {
		return nil, fmt.Errorf("national is not in a bracket")
	}

	return &nationalTeamBracket, nil
}

func (repo NationalTeamBracketRepositoryDb) FindAll() (*[]domain.NationalTeamBracket, error) {
	var nationalTeamsBrackets []domain.NationalTeamBracket

	repo.Db.Find(&nationalTeamsBrackets)

	return &nationalTeamsBrackets, nil
}

func (repo NationalTeamBracketRepositoryDb) Update(nationalTeamBracket *domain.NationalTeamBracket) (*domain.NationalTeamBracket, error) {
	err := repo.Db.Save(&nationalTeamBracket).Error

	if err != nil {
		return nil, err
	}

	return nationalTeamBracket, nil
}
