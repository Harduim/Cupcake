package repositories_test

import (
	"cupcake/app/database"
	"cupcake/app/domain"
	"cupcake/app/repositories"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUserRepositoryDbInsert(t *testing.T) {
	db, err := database.NewTest()
	require.Nil(t, err)

	user := domain.NewUser()
	user.ID = uuid.NewV4().String()
	user.Name = "ANY_NAME"
	user.Email = "any@email.com"

	repo := repositories.UserRepositoryDb{Db: db}
	repo.Insert(user)

	createdUser, err := repo.Find(user.ID)
	require.Equal(t, createdUser.ID, user.ID)
	require.Equal(t, createdUser.Name, user.Name)
	require.Equal(t, createdUser.Email, user.Email)
}

func TestUserRepositoryDbUpdate(t *testing.T) {
	db, err := database.NewTest()
	require.Nil(t, err)

	user := domain.NewUser()
	user.ID = uuid.NewV4().String()
	user.Name = "ANY_NAME"
	user.Email = "any@email.com"

	repo := repositories.UserRepositoryDb{Db: db}
	repo.Insert(user)

	createdUser, err := repo.Find(user.ID)
	require.Equal(t, createdUser.ID, user.ID)
	require.Equal(t, createdUser.Name, user.Name)
	require.Equal(t, createdUser.Email, user.Email)

	user.Name = "ANY_NAME_2"

	updatedUser, err := repo.Update(user)
	require.Equal(t, updatedUser.ID, user.ID)
	require.Equal(t, updatedUser.Name, user.Name)
	require.Equal(t, updatedUser.Email, user.Email)
}
