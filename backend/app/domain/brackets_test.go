package domain_test

import (
	"cupcake/app/domain"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestValidateIfBracketIsEmpty(t *testing.T) {
	bracket := domain.NewBracket()
	err := bracket.Validate()

	require.Error(t, err)
}

func TestBracketIdIsNotAUuid(t *testing.T) {
	bracket := domain.NewBracket()

	bracket.ID = "ANY_ID"
	bracket.Name = "ANY_NAME"
	bracket.Multiplier = 1

	err := bracket.Validate()

	require.Error(t, err)
}

func TestBracketValidation(t *testing.T) {
	bracket := domain.NewBracket()

	bracket.ID = uuid.NewV4().String()
	bracket.Name = "ANY_NAME"
	bracket.Multiplier = 2

	err := bracket.Validate()

	require.Nil(t, err)
}
