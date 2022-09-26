package models_test

import (
	"cupcake/app/models"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestValidateIfNationalTeamIsEmpty(t *testing.T) {
	nationalTeam := models.NewNationalTeam()
	err := nationalTeam.Validate()

	require.Error(t, err)
}

func TestNationalTeamIdIsNotAUuid(t *testing.T) {
	nationalTeam := models.NewNationalTeam()

	nationalTeam.ID = "ANY_ID"
	nationalTeam.Name = "ANY_NAME"
	nationalTeam.KeyID = uuid.NewV4().String()

	err := nationalTeam.Validate()

	require.Error(t, err)
}

func TestNationalTeamKeyIDIsNotAUuid(t *testing.T) {
	nationalTeam := models.NewNationalTeam()

	nationalTeam.ID = uuid.NewV4().String()
	nationalTeam.Name = "ANY_NAME"
	nationalTeam.KeyID = "ANY_KEY"

	err := nationalTeam.Validate()

	require.Error(t, err)
}

func TestNationalTeamValidation(t *testing.T) {
	nationalTeam := models.NewNationalTeam()

	nationalTeam.ID = uuid.NewV4().String()
	nationalTeam.Name = "ANY_NAME"
	nationalTeam.KeyID = uuid.NewV4().String()

	err := nationalTeam.Validate()

	require.Nil(t, err)
}
