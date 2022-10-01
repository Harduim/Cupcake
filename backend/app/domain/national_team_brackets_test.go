package domain_test

import (
	"cupcake/app/domain"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestValidateIfNationalTeamBracketsIsEmpty(t *testing.T) {
	nationalTeam := domain.NewNationalTeamBracket()
	err := nationalTeam.Validate()

	require.Error(t, err)
}

func TestNationalTeamBracketIdIsNotAUuid(t *testing.T) {
	nationalTeam := domain.NewNationalTeamBracket()

	nationalTeam.NationalTeamID = uuid.NewV4().String()
	nationalTeam.BracketID = "ANY_ID"

	err := nationalTeam.Validate()

	require.Error(t, err)
}

func TestNationalTeamBracketNationalTeamIDIsNotAUuid(t *testing.T) {
	nationalTeam := domain.NewNationalTeamBracket()

	nationalTeam.NationalTeamID = "ANY_ID"
	nationalTeam.BracketID = uuid.NewV4().String()

	err := nationalTeam.Validate()

	require.Error(t, err)
}

func TestNationalTeamBracketsValidation(t *testing.T) {
	nationalTeam := domain.NewNationalTeamBracket()

	nationalTeam.NationalTeamID = uuid.NewV4().String()
	nationalTeam.BracketID = uuid.NewV4().String()

	err := nationalTeam.Validate()

	require.Nil(t, err)
}
