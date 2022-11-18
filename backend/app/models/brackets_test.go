package models_test

import (
	domain "cupcake/app/models"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
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
	bracket.OpenDate = time.Date(2022, 12, 10, 15, 0, 0, 0, time.Local)
	bracket.CloseDate = time.Date(2022, 12, 18, 15, 0, 0, 0, time.Local)

	err := bracket.Validate()

	require.Error(t, err)
}

func TestOpenDateIsAfterCloseDate(t *testing.T) {
	bracket := domain.NewBracket()

	bracket.ID = "ANY_ID"
	bracket.Name = "ANY_NAME"
	bracket.Multiplier = 1
	bracket.OpenDate = time.Date(2022, 12, 19, 15, 0, 0, 0, time.Local)
	bracket.CloseDate = time.Date(2022, 12, 18, 15, 0, 0, 0, time.Local)

	err := bracket.Validate()

	require.Error(t, err)
}

func TestBracketValidation(t *testing.T) {
	bracket := domain.NewBracket()

	bracket.ID = uuid.NewV4().String()
	bracket.Name = "ANY_NAME"
	bracket.Multiplier = 2
	bracket.OpenDate = time.Date(2022, 12, 10, 15, 0, 0, 0, time.Local)
	bracket.CloseDate = time.Date(2022, 12, 18, 15, 0, 0, 0, time.Local)

	err := bracket.Validate()

	require.Nil(t, err)
}
