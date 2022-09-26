package models_test

import (
	"cupcake/app/models"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestBetIdIsNotAUuid(t *testing.T) {
	bet := models.NewBet()

	bet.ID = "ANY_ID"
	bet.CreatedAt = time.Time{}
	bet.NationalTeamA = uuid.NewV4().String()
	bet.NationalTeamB = uuid.NewV4().String()
	bet.GolA = 0
	bet.GolB = 0
	bet.UserID = uuid.NewV4().String()
	bet.MatchID = uuid.NewV4().String()

	err := bet.Validate()

	require.Error(t, err)
}

func TestBetNationalTeamAIsNotAUuid(t *testing.T) {
	bet := models.NewBet()

	bet.ID = uuid.NewV4().String()
	bet.CreatedAt = time.Time{}
	bet.NationalTeamA = "ANY_ID"
	bet.NationalTeamB = uuid.NewV4().String()
	bet.GolA = 0
	bet.GolB = 0
	bet.UserID = uuid.NewV4().String()
	bet.MatchID = uuid.NewV4().String()

	err := bet.Validate()

	require.Error(t, err)
}

func TestBetNationalTeamBIsNotAUuid(t *testing.T) {
	bet := models.NewBet()

	bet.ID = uuid.NewV4().String()
	bet.CreatedAt = time.Time{}
	bet.NationalTeamA = uuid.NewV4().String()
	bet.NationalTeamB = "ANY_ID"
	bet.GolA = 0
	bet.GolB = 0
	bet.UserID = uuid.NewV4().String()
	bet.MatchID = uuid.NewV4().String()

	err := bet.Validate()

	require.Error(t, err)
}

func TestBetGolAIsEmpty(t *testing.T) {
	bet := models.NewBet()

	bet.ID = uuid.NewV4().String()
	bet.CreatedAt = time.Time{}
	bet.NationalTeamA = uuid.NewV4().String()
	bet.NationalTeamB = uuid.NewV4().String()
	bet.GolB = 0
	bet.UserID = uuid.NewV4().String()
	bet.MatchID = uuid.NewV4().String()

	err := bet.Validate()

	require.Error(t, err)
}

func TestBetGolBIsEmpty(t *testing.T) {
	bet := models.NewBet()

	bet.ID = uuid.NewV4().String()
	bet.CreatedAt = time.Time{}
	bet.NationalTeamA = uuid.NewV4().String()
	bet.NationalTeamB = uuid.NewV4().String()
	bet.GolA = 0
	bet.UserID = uuid.NewV4().String()
	bet.MatchID = uuid.NewV4().String()

	err := bet.Validate()

	require.Error(t, err)
}

func TestBetUserIDIsNotAUuid(t *testing.T) {
	bet := models.NewBet()

	bet.ID = uuid.NewV4().String()
	bet.CreatedAt = time.Time{}
	bet.NationalTeamA = uuid.NewV4().String()
	bet.NationalTeamB = uuid.NewV4().String()
	bet.GolA = 0
	bet.GolB = 0
	bet.UserID = "ANY_ID"
	bet.MatchID = uuid.NewV4().String()

	err := bet.Validate()

	require.Error(t, err)
}

func TestBetMatchIDIsNotAUuid(t *testing.T) {
	bet := models.NewBet()

	bet.ID = uuid.NewV4().String()
	bet.CreatedAt = time.Time{}
	bet.NationalTeamA = uuid.NewV4().String()
	bet.NationalTeamB = uuid.NewV4().String()
	bet.GolA = 0
	bet.GolB = 0
	bet.UserID = uuid.NewV4().String()
	bet.MatchID = "ANY_ID"

	err := bet.Validate()

	require.Error(t, err)
}

func TestBetDateIsEmpty(t *testing.T) {
	bet := models.NewBet()

	bet.ID = uuid.NewV4().String()
	bet.NationalTeamA = uuid.NewV4().String()
	bet.NationalTeamB = uuid.NewV4().String()
	bet.GolA = 0
	bet.GolB = 0
	bet.UserID = uuid.NewV4().String()
	bet.MatchID = uuid.NewV4().String()

	err := bet.Validate()

	require.Error(t, err)
}
