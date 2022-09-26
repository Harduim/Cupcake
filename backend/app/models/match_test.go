package models_test

import (
	"cupcake/app/models"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestMatchIdIsNotAUuid(t *testing.T) {
	match := models.NewMatch()

	match.ID = "ANY_ID"
	match.Date = time.Time{}
	match.NationalTeamA = uuid.NewV4().String()
	match.NationalTeamB = uuid.NewV4().String()
	match.GolA = 0
	match.GolB = 0
	match.KeyID = uuid.NewV4().String()

	err := match.Validate()

	require.Error(t, err)
}

func TestMatchNationalTeamAIsNotAUuid(t *testing.T) {
	match := models.NewMatch()

	match.ID = uuid.NewV4().String()
	match.Date = time.Time{}
	match.NationalTeamA = "ANY_ID"
	match.NationalTeamB = uuid.NewV4().String()
	match.GolA = 0
	match.GolB = 0
	match.KeyID = uuid.NewV4().String()

	err := match.Validate()

	require.Error(t, err)
}

func TestMatchNationalTeamBIsNotAUuid(t *testing.T) {
	match := models.NewMatch()

	match.ID = uuid.NewV4().String()
	match.Date = time.Time{}
	match.NationalTeamA = uuid.NewV4().String()
	match.NationalTeamB = "ANY_ID"
	match.GolA = 0
	match.GolB = 0
	match.KeyID = uuid.NewV4().String()

	err := match.Validate()

	require.Error(t, err)
}

func TestMatchGolAIsEmpty(t *testing.T) {
	match := models.NewMatch()

	match.ID = uuid.NewV4().String()
	match.Date = time.Time{}
	match.NationalTeamA = uuid.NewV4().String()
	match.NationalTeamB = uuid.NewV4().String()
	match.GolB = 0
	match.KeyID = uuid.NewV4().String()

	err := match.Validate()

	require.Error(t, err)
}

func TestMatchGolBIsEmpty(t *testing.T) {
	match := models.NewMatch()

	match.ID = uuid.NewV4().String()
	match.Date = time.Time{}
	match.NationalTeamA = uuid.NewV4().String()
	match.NationalTeamB = uuid.NewV4().String()
	match.GolA = 0
	match.KeyID = uuid.NewV4().String()

	err := match.Validate()

	require.Error(t, err)
}

func TestMatchKeyIDIsNotAUuid(t *testing.T) {
	match := models.NewMatch()

	match.ID = uuid.NewV4().String()
	match.Date = time.Time{}
	match.NationalTeamA = uuid.NewV4().String()
	match.NationalTeamB = uuid.NewV4().String()
	match.GolA = 0
	match.GolB = 0
	match.KeyID = "ANY_ID"

	err := match.Validate()

	require.Error(t, err)
}

func TestMatchDateIsEmpty(t *testing.T) {
	match := models.NewMatch()

	match.ID = uuid.NewV4().String()
	match.NationalTeamA = uuid.NewV4().String()
	match.NationalTeamB = uuid.NewV4().String()
	match.GolA = 0
	match.GolB = 0
	match.KeyID = uuid.NewV4().String()

	err := match.Validate()

	require.Error(t, err)
}
