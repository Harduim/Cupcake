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

func MakeBet() *domain.Bet {
	bet := domain.NewBet()

	bet.ID = uuid.NewV4().String()
	bet.CreatedAt = time.Now().UTC()
	bet.NationalTeamAID = uuid.NewV4().String()
	bet.NationalTeamBID = uuid.NewV4().String()
	bet.GolA = 0
	bet.GolB = 0
	bet.UserID = uuid.NewV4().String()
	bet.MatchID = uuid.NewV4().String()

	return bet
}

func TestBetRepositoryDbInsert(t *testing.T) {
	db, err := database.NewTest()
	require.Nil(t, err)

	bet := MakeBet()

	repo := repositories.BetRepositoryDb{Db: db}
	_, err = repo.Insert(bet)

	require.Nil(t, err)

	createdBet, err := repo.Find(bet.ID)

	require.Equal(t, createdBet.ID, bet.ID)
	require.Equal(t, createdBet.CreatedAt, bet.CreatedAt)
	require.Equal(t, createdBet.NationalTeamAID, bet.NationalTeamAID)
	require.Equal(t, createdBet.NationalTeamBID, bet.NationalTeamBID)
	require.Equal(t, createdBet.GolA, bet.GolA)
	require.Equal(t, createdBet.GolB, bet.GolB)
	require.Equal(t, createdBet.UserID, bet.UserID)
	require.Equal(t, createdBet.MatchID, bet.MatchID)
}

func TestBetRepositoryDbUpdate(t *testing.T) {
	db, err := database.NewTest()
	require.Nil(t, err)

	bet := MakeBet()

	repo := repositories.BetRepositoryDb{Db: db}
	_, err = repo.Insert(bet)

	if err != nil {
		return
	}

	require.Nil(t, err)

	createdBet, err := repo.Find(bet.ID)

	require.Equal(t, createdBet.ID, bet.ID)
	require.Equal(t, createdBet.CreatedAt, bet.CreatedAt)
	require.Equal(t, createdBet.NationalTeamAID, bet.NationalTeamAID)
	require.Equal(t, createdBet.NationalTeamBID, bet.NationalTeamBID)
	require.Equal(t, createdBet.GolA, bet.GolA)
	require.Equal(t, createdBet.GolB, bet.GolB)
	require.Equal(t, createdBet.UserID, bet.UserID)
	require.Equal(t, createdBet.MatchID, bet.MatchID)

	createdBet.NationalTeamAID = uuid.NewV4().String()
	updatedBet, err := repo.Update(createdBet)

	require.Equal(t, createdBet.ID, updatedBet.ID)
	require.Equal(t, createdBet.CreatedAt, updatedBet.CreatedAt)
	require.Equal(t, createdBet.NationalTeamAID, updatedBet.NationalTeamAID)
	require.Equal(t, createdBet.NationalTeamBID, updatedBet.NationalTeamBID)
	require.Equal(t, createdBet.GolA, updatedBet.GolA)
	require.Equal(t, createdBet.GolB, updatedBet.GolB)
	require.Equal(t, createdBet.UserID, updatedBet.UserID)
	require.Equal(t, createdBet.MatchID, updatedBet.MatchID)
}
