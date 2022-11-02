package domain_test

import (
	"cupcake/app/domain"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestJokerNationalTeamAIsNotAUuid(t *testing.T) {
	nationalTeamAID := "ANY_ID"
	nationalTeamBID := uuid.NewV4().String()
	golA := 0
	golB := 0
	userID := uuid.NewV4().String()
	matchID := uuid.NewV4().String()
	winnerID := uuid.NewV4().String()

	_, err := domain.NewJoker(nationalTeamAID, nationalTeamBID, matchID, userID, &golA, &golB, winnerID)

	require.Error(t, err)
}

func TestJokerNationalTeamBIsNotAUuid(t *testing.T) {
	nationalTeamAID := uuid.NewV4().String()
	nationalTeamBID := "ANY_ID"
	golA := 0
	golB := 0
	userID := uuid.NewV4().String()
	matchID := uuid.NewV4().String()
	winnerID := uuid.NewV4().String()

	_, err := domain.NewJoker(nationalTeamAID, nationalTeamBID, matchID, userID, &golA, &golB, winnerID)

	require.Error(t, err)
}

func TestJokerUserIDIsNotAUuid(t *testing.T) {
	nationalTeamAID := uuid.NewV4().String()
	nationalTeamBID := uuid.NewV4().String()
	golA := 0
	golB := 0
	userID := "ANY_ID"
	matchID := uuid.NewV4().String()
	winnerID := uuid.NewV4().String()

	_, err := domain.NewJoker(nationalTeamAID, nationalTeamBID, matchID, userID, &golA, &golB, winnerID)

	require.Error(t, err)
}

func TestJokerMatchIDIsNotAUuid(t *testing.T) {
	nationalTeamAID := uuid.NewV4().String()
	nationalTeamBID := uuid.NewV4().String()
	golA := 0
	golB := 0
	userID := uuid.NewV4().String()
	matchID := "ANY_ID"
	winnerID := uuid.NewV4().String()

	_, err := domain.NewJoker(nationalTeamAID, nationalTeamBID, matchID, userID, &golA, &golB, winnerID)

	require.Error(t, err)
}

func TestJokerWinnerIDIsNotAUuid(t *testing.T) {
	nationalTeamAID := uuid.NewV4().String()
	nationalTeamBID := uuid.NewV4().String()
	golA := 0
	golB := 0
	userID := uuid.NewV4().String()
	matchID := uuid.NewV4().String()
	winnerID := "ANY_ID"

	_, err := domain.NewJoker(nationalTeamAID, nationalTeamBID, matchID, userID, &golA, &golB, winnerID)

	require.Error(t, err)
}
