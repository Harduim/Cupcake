package repositories

import (
	"cupcake/app/database"
	"cupcake/app/domain"
	"fmt"
)

type NationalTeamRepository interface {
	Insert(nationalTeam *domain.NationalTeam) (*domain.NationalTeam, error)
	Find(id string) (*domain.NationalTeam, error)
	Update(nationalTeam *domain.NationalTeam) (*domain.NationalTeam, error)
}

type NationalTeamRepositoryDb struct {
	Db *database.Database
}

func (repo NationalTeamRepositoryDb) Insert(nationalTeam *domain.NationalTeam) (*domain.NationalTeam, error) {
	err := repo.Db.Create(nationalTeam).Error

	if err != nil {
		return nil, err
	}

	return nationalTeam, nil
}

func (repo NationalTeamRepositoryDb) Find(id string) (*domain.NationalTeam, error) {
	var nationalTeam domain.NationalTeam

	repo.Db.First(&nationalTeam, "id = ?", id)

	if nationalTeam.ID == "" {
		return nil, fmt.Errorf("nationalTeam does not exist")
	}

	return &nationalTeam, nil
}

func (repo NationalTeamRepositoryDb) FindAll() (*[]domain.NationalTeam, error) {
	var nationalTeams []domain.NationalTeam

	repo.Db.Find(&nationalTeams)

	return &nationalTeams, nil
}

func (repo NationalTeamRepositoryDb) Update(nationalTeam *domain.NationalTeam) (*domain.NationalTeam, error) {
	err := repo.Db.Save(&nationalTeam).Error

	if err != nil {
		return nil, err
	}

	return nationalTeam, nil
}
