package repositories_test

import (
	"cupcake/app/database"
	domain "cupcake/app/models"
	"cupcake/app/repositories"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func MakeNationalTeam() *domain.NationalTeam {
	nationalTeam := domain.NewNationalTeam()

	nationalTeam.ID = uuid.NewV4().String()
	nationalTeam.Name = "ANY_NAME"

	return nationalTeam
}

func TestNationalTeamRepositoryDbInsert(t *testing.T) {
	db, err := database.NewTest()
	require.Nil(t, err)

	nationalTeam := MakeNationalTeam()

	repo := repositories.NationalTeamRepositoryDb{Db: db}
	_, err = repo.Insert(nationalTeam)

	require.Nil(t, err)

	createdNationalTeam, err := repo.Find(nationalTeam.ID)

	require.Equal(t, createdNationalTeam.ID, nationalTeam.ID)
	require.Equal(t, createdNationalTeam.Name, nationalTeam.Name)
}

func TestNationalTeamRepositoryDbUpdate(t *testing.T) {
	db, err := database.NewTest()
	require.Nil(t, err)

	nationalTeam := MakeNationalTeam()

	repo := repositories.NationalTeamRepositoryDb{Db: db}
	_, err = repo.Insert(nationalTeam)

	require.Nil(t, err)

	createdNationalTeam, err := repo.Find(nationalTeam.ID)

	require.Equal(t, createdNationalTeam.ID, nationalTeam.ID)
	require.Equal(t, createdNationalTeam.Name, nationalTeam.Name)

	createdNationalTeam.Name = "NEW_NAME"
	updatedNationalTeam, err := repo.Update(createdNationalTeam)

	require.Equal(t, createdNationalTeam.ID, updatedNationalTeam.ID)
	require.Equal(t, createdNationalTeam.Name, updatedNationalTeam.Name)
}
