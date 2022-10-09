package repositories_test

import (
	"cupcake/app/database"
	"cupcake/app/domain"
	"cupcake/app/repositories"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func MakeNationalTeamBracket() *domain.NationalTeamBracket {
	nationalTeamBracket := domain.NewNationalTeamBracket()

	nationalTeamBracket.NationalTeamID = uuid.NewV4().String()
	nationalTeamBracket.BracketID = uuid.NewV4().String()

	return nationalTeamBracket
}

func TestNationalTeamBracketRepositoryDbInsert(t *testing.T) {
	db, err := database.NewTest()
	require.Nil(t, err)

	nationalTeamBracket := MakeNationalTeamBracket()

	repo := repositories.NationalTeamBracketRepositoryDb{Db: db}
	_, err = repo.Insert(nationalTeamBracket)

	require.Nil(t, err)

	createdNationalTeamBracket, err := repo.Find(nationalTeamBracket.NationalTeamID)

	require.Equal(t, createdNationalTeamBracket.BracketID, nationalTeamBracket.BracketID)
	require.Equal(t, createdNationalTeamBracket.NationalTeamID, nationalTeamBracket.NationalTeamID)
}

func TestNationalTeamBracketRepositoryDbUpdate(t *testing.T) {
	db, err := database.NewTest()
	require.Nil(t, err)

	nationalTeamBracket := MakeNationalTeamBracket()

	repo := repositories.NationalTeamBracketRepositoryDb{Db: db}
	_, err = repo.Insert(nationalTeamBracket)

	require.Nil(t, err)

	createdNationalTeamBracket, err := repo.Find(nationalTeamBracket.NationalTeamID)

	require.Equal(t, createdNationalTeamBracket.BracketID, nationalTeamBracket.BracketID)
	require.Equal(t, createdNationalTeamBracket.NationalTeamID, nationalTeamBracket.NationalTeamID)

	createdNationalTeamBracket.BracketID = uuid.NewV4().String()
	updatedNationalTeamBracket, err := repo.Update(createdNationalTeamBracket)

	require.Equal(t, createdNationalTeamBracket.BracketID, updatedNationalTeamBracket.BracketID)
	require.Equal(t, createdNationalTeamBracket.NationalTeamID, updatedNationalTeamBracket.NationalTeamID)
}
