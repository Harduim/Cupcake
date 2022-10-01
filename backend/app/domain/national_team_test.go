package domain_test

import (
	"cupcake/app/domain"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestValidateIfNationalTeamIsEmpty(t *testing.T) {
	nationalTeam := domain.NewNationalTeam()
	err := nationalTeam.Validate()

	require.Error(t, err)
}

func TestNationalTeamIdIsNotAUuid(t *testing.T) {
	nationalTeam := domain.NewNationalTeam()

	nationalTeam.ID = "ANY_ID"
	nationalTeam.Name = "ANY_NAME"

	err := nationalTeam.Validate()

	require.Error(t, err)
}

func TestNationalTeamValidation(t *testing.T) {
	nationalTeam := domain.NewNationalTeam()

	nationalTeam.ID = uuid.NewV4().String()
	nationalTeam.Name = "ANY_NAME"

	err := nationalTeam.Validate()

	require.Nil(t, err)
}
