package domain_test

import (
	domain "cupcake/app/models"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestValidateIfUserIsEmpty(t *testing.T) {
	user := domain.NewUser()
	err := user.Validate()

	require.Error(t, err)
}

func TestUserIdIsNotAUuid(t *testing.T) {
	user := domain.NewUser()
	falseBoolean := false

	user.ID = "ANY_ID"
	user.Name = "ANY_NAME"
	user.Email = "ANY_EMAIL"
	user.IsAdmin = &falseBoolean

	err := user.Validate()

	require.Error(t, err)
}

func TestUserEmailValidation(t *testing.T) {
	user := domain.NewUser()
	falseBoolean := false

	user.ID = uuid.NewV4().String()
	user.Name = "ANY_NAME"
	user.Email = "INVALID_EMAIL"
	user.IsAdmin = &falseBoolean

	err := user.Validate()

	require.Error(t, err)
}

func TestUserValidation(t *testing.T) {
	user := domain.NewUser()
	falseBoolean := false

	user.ID = uuid.NewV4().String()
	user.Name = "ANY_NAME"
	user.Email = "example@test.com"
	user.IsAdmin = &falseBoolean

	err := user.Validate()

	require.Nil(t, err)
}
