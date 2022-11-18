package repositories

import (
	"cupcake/app/database"
	"cupcake/app/models"
	"fmt"
)

type NationalTeamRepository interface {
	Insert(nationalTeam *models.NationalTeam) (*models.NationalTeam, error)
	Find(id string) (*models.NationalTeam, error)
	FindAll() ([]*models.NationalTeam, error)
	Update(nationalTeam *models.NationalTeam) (*models.NationalTeam, error)
}

type NationalTeamRepositoryDb struct {
	Db *database.Database
}

func (repo NationalTeamRepositoryDb) Insert(nationalTeam *models.NationalTeam) (*models.NationalTeam, error) {
	err := repo.Db.Create(nationalTeam).Error

	if err != nil {
		return nil, err
	}

	return nationalTeam, nil
}

func (repo NationalTeamRepositoryDb) Find(id string) (*models.NationalTeam, error) {
	var nationalTeam models.NationalTeam

	repo.Db.First(&nationalTeam, "id = ?", id)

	if nationalTeam.ID == "" {
		return nil, fmt.Errorf("nationalTeam does not exist")
	}

	return &nationalTeam, nil
}

func (repo NationalTeamRepositoryDb) FindAll() (*[]models.NationalTeam, error) {
	var nationalTeams []models.NationalTeam

	repo.Db.Find(&nationalTeams)

	return &nationalTeams, nil
}

func (repo NationalTeamRepositoryDb) Update(nationalTeam *models.NationalTeam) (*models.NationalTeam, error) {
	err := repo.Db.Save(&nationalTeam).Error

	if err != nil {
		return nil, err
	}

	return nationalTeam, nil
}
