package repositories_test

import (
	"cupcake/app/database"
	"cupcake/app/domain"
	"cupcake/app/repositories"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func MakeJoker() *domain.Joker {
	nationalTeamAID := uuid.NewV4().String()
	nationalTeamBID := uuid.NewV4().String()
	golA := 0
	golB := 0
	userID := uuid.NewV4().String()
	bracketID := uuid.NewV4().String()
	winnerID := uuid.NewV4().String()
	joker, _ := domain.NewJoker(nationalTeamAID, nationalTeamBID, bracketID, userID, &golA, &golB, winnerID)
	return joker
}

func TestJokerRepositoryDbInsert(t *testing.T) {
	db, err := database.NewTest()
	require.Nil(t, err)

	joker := MakeJoker()

	repo := repositories.JokerRepositoryDb{Db: db}
	_, err = repo.Insert(joker)

	require.Nil(t, err)

	createdJoker, err := repo.Find(joker.ID)

	require.Equal(t, createdJoker.ID, joker.ID)
	require.Equal(t, createdJoker.CreatedAt, joker.CreatedAt)
	require.Equal(t, createdJoker.NationalTeamAID, joker.NationalTeamAID)
	require.Equal(t, createdJoker.NationalTeamBID, joker.NationalTeamBID)
	require.Equal(t, createdJoker.GolA, joker.GolA)
	require.Equal(t, createdJoker.GolB, joker.GolB)
	require.Equal(t, createdJoker.UserID, joker.UserID)
	require.Equal(t, createdJoker.BracketID, joker.BracketID)
	require.Equal(t, createdJoker.WinnerID, joker.WinnerID)
}

func TestJokerRepositoryDbUpdate(t *testing.T) {
	db, err := database.NewTest()
	require.Nil(t, err)

	joker := MakeJoker()

	repo := repositories.JokerRepositoryDb{Db: db}
	_, err = repo.Insert(joker)

	if err != nil {
		return
	}

	require.Nil(t, err)

	createdJoker, err := repo.Find(joker.ID)

	require.Equal(t, createdJoker.ID, joker.ID)
	require.Equal(t, createdJoker.CreatedAt, joker.CreatedAt)
	require.Equal(t, createdJoker.NationalTeamAID, joker.NationalTeamAID)
	require.Equal(t, createdJoker.NationalTeamBID, joker.NationalTeamBID)
	require.Equal(t, createdJoker.GolA, joker.GolA)
	require.Equal(t, createdJoker.GolB, joker.GolB)
	require.Equal(t, createdJoker.UserID, joker.UserID)
	require.Equal(t, createdJoker.BracketID, joker.BracketID)
	require.Equal(t, createdJoker.WinnerID, joker.WinnerID)

	createdJoker.NationalTeamAID = uuid.NewV4().String()
	updatedJoker, err := repo.Update(createdJoker)

	require.Equal(t, createdJoker.ID, updatedJoker.ID)
	require.Equal(t, createdJoker.CreatedAt, updatedJoker.CreatedAt)
	require.Equal(t, createdJoker.NationalTeamAID, updatedJoker.NationalTeamAID)
	require.Equal(t, createdJoker.NationalTeamBID, updatedJoker.NationalTeamBID)
	require.Equal(t, createdJoker.GolA, updatedJoker.GolA)
	require.Equal(t, createdJoker.GolB, updatedJoker.GolB)
	require.Equal(t, createdJoker.UserID, updatedJoker.UserID)
	require.Equal(t, createdJoker.BracketID, updatedJoker.BracketID)
	require.Equal(t, createdJoker.WinnerID, updatedJoker.WinnerID)
}
