package models_test

import (
	"cupcake/app/models"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestValidateIfKeyIsEmpty(t *testing.T) {
	key := models.NewKey()
	err := key.Validate()

	require.Error(t, err)
}

func TestKeyIdIsNotAUuid(t *testing.T) {
	key := models.NewKey()

	key.ID = "ANY_ID"
	key.Name = "ANY_NAME"

	err := key.Validate()

	require.Error(t, err)
}

func TestKeyValidation(t *testing.T) {
	key := models.NewKey()

	key.ID = uuid.NewV4().String()
	key.Name = "ANY_NAME"

	err := key.Validate()

	require.Nil(t, err)
}
