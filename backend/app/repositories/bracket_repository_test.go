package repositories_test

import (
	"cupcake/app/database"
	domain "cupcake/app/models"
	"cupcake/app/repositories"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func MakeBracket() *domain.Bracket {
	bracket := domain.NewBracket()

	bracket.ID = uuid.NewV4().String()
	bracket.Name = uuid.NewV4().String()
	bracket.Multiplier = 1

	return bracket
}

func TestBracketRepositoryDbInsert(t *testing.T) {
	db, err := database.NewTest()
	require.Nil(t, err)

	bracket := MakeBracket()

	repo := repositories.BracketRepositoryDb{Db: db}
	_, err = repo.Insert(bracket)

	require.Nil(t, err)

	createdBracket, err := repo.Find(bracket.ID)

	require.Equal(t, createdBracket.ID, bracket.ID)
	require.Equal(t, createdBracket.Name, bracket.Name)
	require.Equal(t, createdBracket.Multiplier, bracket.Multiplier)
}

func TestBracketRepositoryDbUpdate(t *testing.T) {
	db, err := database.NewTest()
	require.Nil(t, err)

	bracket := MakeBracket()

	repo := repositories.BracketRepositoryDb{Db: db}
	_, err = repo.Insert(bracket)

	require.Nil(t, err)

	createdBracket, err := repo.Find(bracket.ID)

	require.Equal(t, createdBracket.ID, bracket.ID)
	require.Equal(t, createdBracket.Name, bracket.Name)
	require.Equal(t, createdBracket.Multiplier, bracket.Multiplier)

	createdBracket.Name = "NEW_NAME"
	updatedBracket, err := repo.Update(createdBracket)

	require.Equal(t, createdBracket.ID, updatedBracket.ID)
	require.Equal(t, createdBracket.Name, updatedBracket.Name)
	require.Equal(t, createdBracket.Multiplier, updatedBracket.Multiplier)
}
