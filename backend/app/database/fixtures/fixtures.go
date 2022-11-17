package fixtures

import (
	Config "cupcake/app/config"
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

	err = groupsFixtures(db)

	if err != nil {
		return err
	}

	return nil
}

func bracketFixtures(db *database.Database) error {
	repo := repositories.BracketRepositoryDb{Db: db}

	final := domain.Bracket{
		ID:         Config.BRKT_FINAIS,
		Name:       "Final",
		Multiplier: 8,
		OpenDate:   time.Date(2022, 12, 10, 0, 0, 0, 0, time.Local),
		CloseDate:  time.Date(2022, 12, 18, 14, 0, 0, 0, time.Local),
	}

	terceiro := domain.Bracket{
		ID:         Config.BRKT_TERCERIRO,
		Name:       "Disputa pelo 3º lugar",
		Multiplier: 3,
		OpenDate:   time.Date(2022, 12, 12, 0, 0, 0, 0, time.Local),
		CloseDate:  time.Date(2022, 12, 17, 14, 0, 0, 0, time.Local),
	}

	semi := domain.Bracket{
		ID:         Config.BRKT_SEMIFINAIS,
		Name:       "Semi Final",
		Multiplier: 3,
		OpenDate:   time.Date(2022, 12, 4, 0, 0, 0, 0, time.Local),
		CloseDate:  time.Date(2022, 12, 14, 18, 0, 0, 0, time.Local),
	}

	quartas := domain.Bracket{
		ID:         Config.BRKT_QUARTAS,
		Name:       "Quartas",
		Multiplier: 2,
		OpenDate:   time.Date(2022, 12, 2, 0, 0, 0, 0, time.Local),
		CloseDate:  time.Date(2022, 12, 10, 12, 0, 0, 0, time.Local),
	}

	oitavas := domain.Bracket{
		ID:         Config.BRKT_OITAVAS,
		Name:       "Oitavas",
		Multiplier: 1,
		OpenDate:   time.Date(2022, 11, 30, 12, 0, 0, 0, time.Local),
		CloseDate:  time.Date(2022, 12, 6, 18, 0, 0, 0, time.Local),
	}

	coringa := domain.Bracket{
		ID:         Config.BRKT_CORINGA,
		Name:       "Rodada Coringa",
		Multiplier: 9,
		OpenDate:   time.Date(2022, 11, 10, 0, 0, 0, 0, time.Local),
		CloseDate:  time.Date(2022, 12, 6, 0, 0, 0, 0, time.Local),
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

	_, err = repo.Insert(&coringa)

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
		ID:       Config.NT_BRAZIL,
		Name:     "Brazil",
		Brackets: []domain.Bracket{bracket},
	}

	_, err := repo.Insert(&brazil)
	if err != nil {
		return err
	}

	france := domain.NationalTeam{
		ID:       Config.NT_FRANCE,
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
	matches := []domain.Match{
		// Coringa
		{
			ID:        Config.MATCH_CORINGA,
			Date:      time.Date(2022, 12, 18, 15, 0, 0, 0, time.Local),
			BracketID: Config.BRKT_CORINGA,
		},
		// Oitavas
		{
			ID:        Config.MATCH_OITAVAS_01,
			Date:      time.Date(2022, 12, 18, 15, 0, 0, 0, time.Local),
			BracketID: Config.BRKT_OITAVAS,
		},
		{
			ID:        Config.MATCH_OITAVAS_02,
			Date:      time.Date(2022, 12, 18, 15, 0, 0, 0, time.Local),
			BracketID: Config.BRKT_OITAVAS,
		},
		{
			ID:        Config.MATCH_OITAVAS_03,
			Date:      time.Date(2022, 12, 18, 15, 0, 0, 0, time.Local),
			BracketID: Config.BRKT_OITAVAS,
		},
		{
			ID:        Config.MATCH_OITAVAS_04,
			Date:      time.Date(2022, 12, 18, 15, 0, 0, 0, time.Local),
			BracketID: Config.BRKT_OITAVAS,
		},
		{
			ID:        Config.MATCH_OITAVAS_05,
			Date:      time.Date(2022, 12, 18, 15, 0, 0, 0, time.Local),
			BracketID: Config.BRKT_OITAVAS,
		},
		{
			ID:        Config.MATCH_OITAVAS_06,
			Date:      time.Date(2022, 12, 18, 15, 0, 0, 0, time.Local),
			BracketID: Config.BRKT_OITAVAS,
		},
		{
			ID:        Config.MATCH_OITAVAS_07,
			Date:      time.Date(2022, 12, 18, 15, 0, 0, 0, time.Local),
			BracketID: Config.BRKT_OITAVAS,
		},
		{
			ID:        Config.MATCH_OITAVAS_08,
			Date:      time.Date(2022, 12, 18, 15, 0, 0, 0, time.Local),
			BracketID: Config.BRKT_OITAVAS,
		},
		// Quartas
		{
			ID:        Config.MATCH_QUARTAS_01,
			Date:      time.Date(2022, 12, 18, 15, 0, 0, 0, time.Local),
			BracketID: Config.BRKT_QUARTAS,
		},
		{
			ID:        Config.MATCH_QUARTAS_02,
			Date:      time.Date(2022, 12, 18, 15, 0, 0, 0, time.Local),
			BracketID: Config.BRKT_QUARTAS,
		},
		{
			ID:        Config.MATCH_QUARTAS_03,
			Date:      time.Date(2022, 12, 18, 15, 0, 0, 0, time.Local),
			BracketID: Config.BRKT_QUARTAS,
		},
		{
			ID:        Config.MATCH_QUARTAS_04,
			Date:      time.Date(2022, 12, 18, 15, 0, 0, 0, time.Local),
			BracketID: Config.BRKT_QUARTAS,
		},
		// Semifinais
		{
			ID:        Config.MATCH_SEMIFINAIS_01,
			Date:      time.Date(2022, 12, 18, 15, 0, 0, 0, time.Local),
			BracketID: Config.BRKT_SEMIFINAIS,
		},
		{
			ID:        Config.MATCH_SEMIFINAIS_02,
			Date:      time.Date(2022, 12, 18, 15, 0, 0, 0, time.Local),
			BracketID: Config.BRKT_SEMIFINAIS,
		},
	}

	err := db.Create(&matches).Error

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
		MatchID:         Config.MATCH_CORINGA,
		NationalTeamAID: "6d71278a-4eca-42a8-8ec2-fa51a31ef95c", // Brazil
		NationalTeamBID: Config.NT_FRANCE,                       // France
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

func groupsFixtures(db *database.Database) error {
	repo := repositories.GroupsRepositoryDb{Db: db}

	groups := domain.Groups{
		CreatedAt: time.Date(2022, 12, 18, 13, 0, 0, 0, time.Local),
		UserID:    "b8ee5ddd-1137-45de-9071-20e08ba3f51f",
		BracketID: "ef13e77f-b345-4f4d-b4a7-2d1cfb12fa48",
		NationalTeams: []*domain.NationalTeam{{ID: "6d71278a-4eca-42a8-8ec2-fa51a31ef95c"},
			{ID: Config.NT_FRANCE}},
	}
	_, err := repo.Insert(&groups)

	if err != nil {
		return err
	}

	return nil
}
