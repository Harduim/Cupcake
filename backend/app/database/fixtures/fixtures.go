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
		ID: Config.BRKT_OITAVAS,
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
			Date:      time.Date(2022, 12, 6, 0, 0, 0, 0, time.Local),
			BracketID: Config.BRKT_CORINGA,
		},
		// Oitavas
		{
			ID:        Config.MATCH_OITAVAS_01,
			Date:      time.Date(2022, 12, 3, 12, 0, 0, 0, time.Local),
			BracketID: Config.BRKT_OITAVAS,
		},
		{
			ID:        Config.MATCH_OITAVAS_02,
			Date:      time.Date(2022, 12, 3, 16, 0, 0, 0, time.Local),
			BracketID: Config.BRKT_OITAVAS,
		},
		{
			ID:        Config.MATCH_OITAVAS_03,
			Date:      time.Date(2022, 12, 4, 12, 0, 0, 0, time.Local),
			BracketID: Config.BRKT_OITAVAS,
		},
		{
			ID:        Config.MATCH_OITAVAS_04,
			Date:      time.Date(2022, 12, 4, 16, 0, 0, 0, time.Local),
			BracketID: Config.BRKT_OITAVAS,
		},
		{
			ID:        Config.MATCH_OITAVAS_05,
			Date:      time.Date(2022, 12, 5, 12, 0, 0, 0, time.Local),
			BracketID: Config.BRKT_OITAVAS,
		},
		{
			ID:        Config.MATCH_OITAVAS_06,
			Date:      time.Date(2022, 12, 5, 16, 0, 0, 0, time.Local),
			BracketID: Config.BRKT_OITAVAS,
		},
		{
			ID:        Config.MATCH_OITAVAS_07,
			Date:      time.Date(2022, 12, 6, 12, 0, 0, 0, time.Local),
			BracketID: Config.BRKT_OITAVAS,
		},
		{
			ID:        Config.MATCH_OITAVAS_08,
			Date:      time.Date(2022, 12, 6, 16, 0, 0, 0, time.Local),
			BracketID: Config.BRKT_OITAVAS,
		},
		// Quartas
		{
			ID:        Config.MATCH_QUARTAS_01,
			Date:      time.Date(2022, 12, 9, 12, 0, 0, 0, time.Local),
			BracketID: Config.BRKT_QUARTAS,
		},
		{
			ID:        Config.MATCH_QUARTAS_02,
			Date:      time.Date(2022, 12, 9, 16, 0, 0, 0, time.Local),
			BracketID: Config.BRKT_QUARTAS,
		},
		{
			ID:        Config.MATCH_QUARTAS_03,
			Date:      time.Date(2022, 12, 10, 12, 0, 0, 0, time.Local),
			BracketID: Config.BRKT_QUARTAS,
		},
		{
			ID:        Config.MATCH_QUARTAS_04,
			Date:      time.Date(2022, 12, 10, 16, 0, 0, 0, time.Local),
			BracketID: Config.BRKT_QUARTAS,
		},
		// Semifinais
		{
			ID:        Config.MATCH_SEMIFINAIS_01,
			Date:      time.Date(2022, 12, 13, 16, 0, 0, 0, time.Local),
			BracketID: Config.BRKT_SEMIFINAIS,
		},
		{
			ID:        Config.MATCH_SEMIFINAIS_02,
			Date:      time.Date(2022, 12, 14, 16, 0, 0, 0, time.Local),
			BracketID: Config.BRKT_SEMIFINAIS,
		},
		// Terceiro
		{
			ID:        Config.MATCH_TERCEIRO,
			Date:      time.Date(2022, 12, 17, 12, 0, 0, 0, time.Local),
			BracketID: Config.BRKT_TERCERIRO,
		},
		// Finais
		{
			ID:        Config.MATCH_FINAIS,
			Date:      time.Date(2022, 12, 18, 12, 0, 0, 0, time.Local),
			BracketID: Config.BRKT_FINAIS,
		},
	}

	err := db.Create(&matches).Error

	if err != nil {
		return err
	}

	return nil
}
