package models_test

import (
	domain "cupcake/app/models"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestUserPointsValidation(t *testing.T) {
	userPoints := domain.NewUserPoints()

	userPoints.UserID = uuid.NewV4().String()
	userPoints.Points = 1

	err := userPoints.Validate()
	require.Nil(t, err)
}
