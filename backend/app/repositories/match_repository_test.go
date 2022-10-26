package repositories_test

import (
	"cupcake/app/database"
	"cupcake/app/domain"
	"cupcake/app/repositories"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func MakeMatch() *domain.Match {
	match := domain.NewMatch()

	match.ID = uuid.NewV4().String()
	match.Date = time.Time{}
	match.NationalTeamAID = uuid.NewV4().String()
	match.NationalTeamBID = uuid.NewV4().String()
	match.GolA = 0
	match.GolB = 0
	match.BracketID = uuid.NewV4().String()
	match.WinnerID = uuid.NewV4().String()

	return match
}

func TestMatchRepositoryDbInsert(t *testing.T) {
	db, err := database.NewTest()
	require.Nil(t, err)

	match := MakeMatch()

	repo := repositories.MatchRepositoryDb{Db: db}
	repo.Insert(match)

	createdMatch, err := repo.Find(match.ID)

	require.Equal(t, createdMatch.ID, match.ID)
	require.Equal(t, createdMatch.Date, match.Date)
	require.Equal(t, createdMatch.NationalTeamAID, match.NationalTeamAID)
	require.Equal(t, createdMatch.NationalTeamBID, match.NationalTeamBID)
	require.Equal(t, createdMatch.GolA, match.GolA)
	require.Equal(t, createdMatch.GolB, match.GolB)
	require.Equal(t, createdMatch.BracketID, match.BracketID)
	require.Equal(t, createdMatch.WinnerID, match.WinnerID)
}

func TestMatchRepositoryDbUpdate(t *testing.T) {
	db, err := database.NewTest()
	require.Nil(t, err)

	match := MakeMatch()

	repo := repositories.MatchRepositoryDb{Db: db}
	repo.Insert(match)

	createdMatch, err := repo.Find(match.ID)

	require.Equal(t, createdMatch.ID, match.ID)
	require.Equal(t, createdMatch.Date, match.Date)
	require.Equal(t, createdMatch.NationalTeamAID, match.NationalTeamAID)
	require.Equal(t, createdMatch.NationalTeamBID, match.NationalTeamBID)
	require.Equal(t, createdMatch.GolA, match.GolA)
	require.Equal(t, createdMatch.GolB, match.GolB)
	require.Equal(t, createdMatch.BracketID, match.BracketID)
	require.Equal(t, createdMatch.WinnerID, match.WinnerID)

	createdMatch.NationalTeamAID = uuid.NewV4().String()
	updatedMatch, err := repo.Update(createdMatch)

	require.Equal(t, createdMatch.ID, updatedMatch.ID)
	require.Equal(t, createdMatch.Date, updatedMatch.Date)
	require.Equal(t, createdMatch.NationalTeamAID, updatedMatch.NationalTeamAID)
	require.Equal(t, createdMatch.NationalTeamBID, updatedMatch.NationalTeamBID)
	require.Equal(t, createdMatch.GolA, updatedMatch.GolA)
	require.Equal(t, createdMatch.GolB, updatedMatch.GolB)
	require.Equal(t, createdMatch.BracketID, updatedMatch.BracketID)
	require.Equal(t, createdMatch.WinnerID, updatedMatch.WinnerID)
}
