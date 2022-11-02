package domain_test

import (
	"cupcake/app/domain"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBetNationalTeamAIsNotAUuid(t *testing.T) {
	nationalTeamAID := "ANY_ID"
	nationalTeamBID := uuid.NewV4().String()
	golA := 0
	golB := 0
	userID := uuid.NewV4().String()
	matchID := uuid.NewV4().String()
	winnerID := uuid.NewV4().String()

	_, err := domain.NewBet(nationalTeamAID, nationalTeamBID, matchID, userID, &golA, &golB, winnerID)

	require.Error(t, err)
}

func TestBetNationalTeamBIsNotAUuid(t *testing.T) {
	nationalTeamAID := uuid.NewV4().String()
	nationalTeamBID := "ANY_ID"
	golA := 0
	golB := 0
	userID := uuid.NewV4().String()
	matchID := uuid.NewV4().String()
	winnerID := uuid.NewV4().String()

	_, err := domain.NewBet(nationalTeamAID, nationalTeamBID, matchID, userID, &golA, &golB, winnerID)

	require.Error(t, err)
}

func TestBetUserIDIsNotAUuid(t *testing.T) {
	nationalTeamAID := uuid.NewV4().String()
	nationalTeamBID := uuid.NewV4().String()
	golA := 0
	golB := 0
	userID := "ANY_ID"
	matchID := uuid.NewV4().String()
	winnerID := uuid.NewV4().String()

	_, err := domain.NewBet(nationalTeamAID, nationalTeamBID, matchID, userID, &golA, &golB, winnerID)

	require.Error(t, err)
}

func TestBetMatchIDIsNotAUuid(t *testing.T) {
	nationalTeamAID := uuid.NewV4().String()
	nationalTeamBID := uuid.NewV4().String()
	golA := 0
	golB := 0
	userID := uuid.NewV4().String()
	matchID := "ANY_ID"
	winnerID := uuid.NewV4().String()

	_, err := domain.NewBet(nationalTeamAID, nationalTeamBID, matchID, userID, &golA, &golB, winnerID)

	require.Error(t, err)
}

func TestBetWinnerIDIsNotAUuid(t *testing.T) {
	nationalTeamAID := uuid.NewV4().String()
	nationalTeamBID := uuid.NewV4().String()
	golA := 0
	golB := 0
	userID := uuid.NewV4().String()
	matchID := uuid.NewV4().String()
	winnerID := "ANY_ID"

	_, err := domain.NewBet(nationalTeamAID, nationalTeamBID, matchID, userID, &golA, &golB, winnerID)

	require.Error(t, err)
}
