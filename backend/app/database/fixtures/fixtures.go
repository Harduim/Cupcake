package fixtures

import (
	"cupcake/app/database"
	"cupcake/app/domain"
	"cupcake/app/repositories"
	"time"
)

func CreateFixtures(db *database.Database) error {
	err := nationalTeamsFixtures(db)

	if err != nil {
		return err
	}

	err = bracketFixtures(db)

	if err != nil {
		return err
	}

	err = matchFixtures(db)

	if err != nil {
		return err
	}

	err = nationalTeamBracketsFixtures(db)

	if err != nil {
		return err
	}

	err = userFixtures(db)

	if err != nil {
		return err
	}

	err = userPointsFixtures(db)

	if err != nil {
		return err
	}

	err = betFixtures(db)

	if err != nil {
		return err
	}

	return nil
}

func nationalTeamsFixtures(db *database.Database) error {
	repo := repositories.NationalTeamRepositoryDb{Db: db}

	brazil := domain.NationalTeam{
		ID:   "6d71278a-4eca-42a8-8ec2-fa51a31ef95c",
		Name: "Brazil",
	}

	_, err := repo.Insert(&brazil)
	if err != nil {
		return err
	}

	france := domain.NationalTeam{
		ID:   "4935b4e1-f422-41a7-9a22-051f429ff5e4",
		Name: "France",
	}

	_, err = repo.Insert(&france)
	if err != nil {
		return err
	}
	return nil
}

func bracketFixtures(db *database.Database) error {
	repo := repositories.BracketRepositoryDb{Db: db}

	bracket := domain.Bracket{
		ID:         "5ef28a89-f697-4af2-931d-808c41cbd2d1",
		Name:       "Final",
		Multiplier: 3,
	}

	_, err := repo.Insert(&bracket)

	if err != nil {
		return err
	}

	return nil
}

func matchFixtures(db *database.Database) error {
	repo := repositories.MatchRepositoryDb{Db: db}

	match := domain.Match{
		ID:              "719cf785-0753-4864-a0c9-546d1c8cf998",
		Date:            time.Date(2022, 12, 18, 15, 0, 0, 0, time.Local),
		NationalTeamAID: "6d71278a-4eca-42a8-8ec2-fa51a31ef95c", // Brazil
		NationalTeamBID: "4935b4e1-f422-41a7-9a22-051f429ff5e4", // France
		GolA:            2,
		GolB:            0,
		BracketID:       "5ef28a89-f697-4af2-931d-808c41cbd2d1", // Finals
		WinnerID:        "6d71278a-4eca-42a8-8ec2-fa51a31ef95c", // Brazil, of course.
	}

	_, err := repo.Insert(&match)

	if err != nil {
		return err
	}

	return nil
}

func nationalTeamBracketsFixtures(db *database.Database) error {
	repo := repositories.NationalTeamBracketRepositoryDb{Db: db}

	nationalTeamBracketsBrazilFinal := domain.NationalTeamBracket{
		NationalTeamID: "6d71278a-4eca-42a8-8ec2-fa51a31ef95c", // Brazil
		BracketID:      "5ef28a89-f697-4af2-931d-808c41cbd2d1", // Finals
	}

	_, err := repo.Insert(&nationalTeamBracketsBrazilFinal)

	if err != nil {
		return err
	}

	nationalTeamBracketsFranceFinal := domain.NationalTeamBracket{
		NationalTeamID: "4935b4e1-f422-41a7-9a22-051f429ff5e4", // France
		BracketID:      "5ef28a89-f697-4af2-931d-808c41cbd2d1", // Finals
	}

	_, err = repo.Insert(&nationalTeamBracketsFranceFinal)

	if err != nil {
		return err
	}

	return nil
}

func betFixtures(db *database.Database) error {
	repo := repositories.BetRepositoryDb{Db: db}

	bet := domain.Bet{
		ID:              "da3db00f-dfc3-4b83-ad31-0dd78b32ae56",
		CreatedAt:       time.Date(2022, 12, 18, 13, 0, 0, 0, time.Local),
		GolA:            2,
		GolB:            0,
		UserID:          "b8ee5ddd-1137-45de-9071-20e08ba3f51f",
		MatchID:         "719cf785-0753-4864-a0c9-546d1c8cf998",
		NationalTeamAID: "6d71278a-4eca-42a8-8ec2-fa51a31ef95c", // Brazil
		NationalTeamBID: "4935b4e1-f422-41a7-9a22-051f429ff5e4", // France
	}

	_, err := repo.Insert(&bet)

	if err != nil {
		return err
	}

	return nil
}

func userFixtures(db *database.Database) error {
	repo := repositories.UserRepositoryDb{Db: db}
	falseBoolean := false

	user := domain.User{
		ID:      "b8ee5ddd-1137-45de-9071-20e08ba3f51f",
		Name:    "Adriano do TI",
		Email:   "adrianodoti@rioenergy.com",
		IsAdmin: &falseBoolean,
	}

	_, err := repo.Insert(&user)

	if err != nil {
		return err
	}

	return nil
}

func userPointsFixtures(db *database.Database) error {
	repo := repositories.UserPointsRepositoryDb{Db: db}

	userPoints := domain.UserPoints{
		UserID: "b8ee5ddd-1137-45de-9071-20e08ba3f51f",
		Points: 3,
	}

	_, err := repo.Insert(&userPoints)

	if err != nil {
		return err
	}

	return nil
}
