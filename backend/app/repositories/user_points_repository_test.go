package repositories_test

import (
	"cupcake/app/database"
	"cupcake/app/domain"
	"cupcake/app/repositories"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func MakeUserPoints() *domain.UserPoints {
	userPoints := domain.NewUserPoints()

	userPoints.UserID = uuid.NewV4().String()
	userPoints.Points = 1

	return userPoints
}

func TestUserPointsRepositoryDbInsert(t *testing.T) {
	db, err := database.NewTest()
	require.Nil(t, err)

	userPoints := MakeUserPoints()

	repo := repositories.UserPointsRepositoryDb{Db: db}
	_, err = repo.Insert(userPoints)

	require.Nil(t, err)

	createdUserPoints, err := repo.Find(userPoints.UserID)

	require.Equal(t, createdUserPoints.UserID, userPoints.UserID)
	require.Equal(t, createdUserPoints.Points, userPoints.Points)
}

func TestUserPointsRepositoryDbUpdate(t *testing.T) {
	db, err := database.NewTest()
	require.Nil(t, err)

	userPoints := MakeUserPoints()

	repo := repositories.UserPointsRepositoryDb{Db: db}

	_, err = repo.Insert(userPoints)

	require.Nil(t, err)

	createdUserPoints, err := repo.Find(userPoints.UserID)

	require.Equal(t, createdUserPoints.UserID, userPoints.UserID)
	require.Equal(t, createdUserPoints.Points, userPoints.Points)

	createdUserPoints.Points = 2
	updatedUserPoints, err := repo.Update(createdUserPoints)

	require.Nil(t, err)

	require.Equal(t, createdUserPoints.UserID, updatedUserPoints.UserID)
	require.Equal(t, createdUserPoints.Points, updatedUserPoints.Points)
}
