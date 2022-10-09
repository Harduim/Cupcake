package domain_test

import (
	"cupcake/app/domain"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUserPointsValidation(t *testing.T) {
	userPoints := domain.NewUserPoints()

	userPoints.UserID = uuid.NewV4().String()
	userPoints.Points = 1

	err := userPoints.Validate()
	require.Nil(t, err)
}
