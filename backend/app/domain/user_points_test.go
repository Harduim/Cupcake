package domain_test

import (
	"cupcake/app/domain"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func MakeNewUser() *domain.User {
	user := domain.NewUser()
	falseBoolean := false

	user.ID = uuid.NewV4().String()
	user.Name = "ANY_NAME"
	user.Email = "example@test.com"
	user.IsAdmin = &falseBoolean

	return user
}

func TestUserPointsValidation(t *testing.T) {
	user := MakeNewUser()

	userPoints := domain.NewUserPoints()
	userPoints.User = user
	userPoints.Points = 0

	err := userPoints.Validate()
	require.Nil(t, err)
}
