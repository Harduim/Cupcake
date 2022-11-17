package fixtures

import (
	"cupcake/app/database"
	"cupcake/app/domain"
	"cupcake/app/repositories"
	"time"
)

func CreateFixtures(db *database.Database) error {
	err := bracketFixtures(db)

	if err != nil {
		return err
	}

	err = nationalTeamsFixtures(db)

	if err != nil {
		return err
	}

	err = matchFixtures(db)

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

	err = jokerFixtures(db)

	if err != nil {
		return err
	}

	err = groupsFixtures(db)

	if err != nil {
		return err
	}

	return nil
}

func bracketFixtures(db *database.Database) error {
	repo := repositories.BracketRepositoryDb{Db: db}

	final := domain.Bracket{
		ID:         "5ef28a89-f697-4af2-931d-808c41cbd2d1",
		Name:       "Final",
		Multiplier: 8,
		OpenDate:   time.Date(2022, 12, 10, 0, 0, 0, 0, time.Local),
		CloseDate:  time.Date(2022, 12, 18, 14, 0, 0, 0, time.Local),
	}

	terceiro := domain.Bracket{
		ID:         "5e87f7d4-ac15-4f8b-b82c-3f50f2d5371f",
		Name:       "Disputa pelo 3º lugar",
		Multiplier: 3,
		OpenDate:   time.Date(2022, 12, 12, 0, 0, 0, 0, time.Local),
		CloseDate:  time.Date(2022, 12, 17, 14, 0, 0, 0, time.Local),
	}

	semi := domain.Bracket{
		ID:         "40e58268-0fc5-4dec-8fcb-b52b46006215",
		Name:       "Semi Final",
		Multiplier: 3,
		OpenDate:   time.Date(2022, 12, 4, 0, 0, 0, 0, time.Local),
		CloseDate:  time.Date(2022, 12, 14, 18, 0, 0, 0, time.Local),
	}

	quartas := domain.Bracket{
		ID:         "22ecea42-848e-43d1-a387-5de1bd468338",
		Name:       "Quartas",
		Multiplier: 2,
		OpenDate:   time.Date(2022, 12, 2, 0, 0, 0, 0, time.Local),
		CloseDate:  time.Date(2022, 12, 10, 12, 0, 0, 0, time.Local),
	}

	oitavas := domain.Bracket{
		ID:         "ef13e77f-b345-4f4d-b4a7-2d1cfb12fa48",
		Name:       "Oitavas",
		Multiplier: 1,
		OpenDate:   time.Date(2022, 11, 30, 12, 0, 0, 0, time.Local),
		CloseDate:  time.Date(2022, 12, 6, 18, 0, 0, 0, time.Local),
	}

	_, err := repo.Insert(&final)

	if err != nil {
		return err
	}

	_, err = repo.Insert(&oitavas)

	if err != nil {
		return err
	}

	_, err = repo.Insert(&terceiro)

	if err != nil {
		return err
	}

	_, err = repo.Insert(&quartas)

	if err != nil {
		return err
	}

	_, err = repo.Insert(&semi)

	if err != nil {
		return err
	}

	return nil
}

func nationalTeamsFixtures(db *database.Database) error {
	repo := repositories.NationalTeamRepositoryDb{Db: db}
	bracket := domain.Bracket{
		ID: "5ef28a89-f697-4af2-931d-808c41cbd2d1",
	}
	brazil := domain.NationalTeam{
		ID:       "6d71278a-4eca-42a8-8ec2-fa51a31ef95c",
		Name:     "Brazil",
		Brackets: []domain.Bracket{bracket},
	}

	_, err := repo.Insert(&brazil)
	if err != nil {
		return err
	}

	france := domain.NationalTeam{
		ID:       "4935b4e1-f422-41a7-9a22-051f429ff5e4",
		Name:     "France",
		Brackets: []domain.Bracket{bracket},
	}

	_, err = repo.Insert(&france)
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

func betFixtures(db *database.Database) error {
	repo := repositories.BetRepositoryDb{Db: db}
	golA := 3
	golB := 2
	bet := domain.Bet{
		ID:              "da3db00f-dfc3-4b83-ad31-0dd78b32ae56",
		CreatedAt:       time.Date(2022, 12, 18, 13, 0, 0, 0, time.Local),
		GolA:            &golA,
		GolB:            &golB,
		UserID:          "b8ee5ddd-1137-45de-9071-20e08ba3f51f",
		MatchID:         "719cf785-0753-4864-a0c9-546d1c8cf998",
		NationalTeamAID: "6d71278a-4eca-42a8-8ec2-fa51a31ef95c", // Brazil
		NationalTeamBID: "4935b4e1-f422-41a7-9a22-051f429ff5e4", // France
		WinnerID:        "6d71278a-4eca-42a8-8ec2-fa51a31ef95c",
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

func jokerFixtures(db *database.Database) error {
	repo := repositories.JokerRepositoryDb{Db: db}

	golA := 3
	golB := 2
	joker := domain.Joker{
		ID:              "9d28a01a-fc22-4d60-9a2a-33df213b71a3",
		CreatedAt:       time.Date(2022, 12, 18, 13, 0, 0, 0, time.Local),
		GolA:            &golA,
		GolB:            &golB,
		UserID:          "b8ee5ddd-1137-45de-9071-20e08ba3f51f",
		BracketID:       "5ef28a89-f697-4af2-931d-808c41cbd2d1",
		NationalTeamAID: "6d71278a-4eca-42a8-8ec2-fa51a31ef95c", // Brazil
		NationalTeamBID: "4935b4e1-f422-41a7-9a22-051f429ff5e4", // France
		WinnerID:        "6d71278a-4eca-42a8-8ec2-fa51a31ef95c",
	}
	_, err := repo.Insert(&joker)

	if err != nil {
		return err
	}

	return nil
}

func groupsFixtures(db *database.Database) error {
	repo := repositories.GroupsRepositoryDb{Db: db}

	groups := domain.Groups{
		CreatedAt: time.Date(2022, 12, 18, 13, 0, 0, 0, time.Local),
		UserID:    "b8ee5ddd-1137-45de-9071-20e08ba3f51f",
		BracketID: "ef13e77f-b345-4f4d-b4a7-2d1cfb12fa48",
		NationalTeams: []*domain.NationalTeam{{ID: "6d71278a-4eca-42a8-8ec2-fa51a31ef95c"},
			{ID: "4935b4e1-f422-41a7-9a22-051f429ff5e4"}},
	}
	_, err := repo.Insert(&groups)

	if err != nil {
		return err
	}

	return nil
}
