package models_test

import (
	"cupcake/app/models"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func MakeNewUser() *models.User {
	user := models.NewUser()
	falseBoolean := false

	user.ID = uuid.NewV4().String()
	user.Name = "ANY_NAME"
	user.Email = "example@test.com"
	user.IsAdmin = &falseBoolean

	return user
}

func TestUserPointsValidation(t *testing.T) {
	user := MakeNewUser()

	userPoints := models.NewUserPoints()
	userPoints.User = *user
	userPoints.Points = 0

	err := userPoints.Validate()
	require.Nil(t, err)
}
