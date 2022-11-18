package fixtures

import (
	"cupcake/app/config"
	"cupcake/app/database"
	"cupcake/app/models"
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
	final := models.Bracket{
		ID:         config.BRKT_FINAIS,
		Name:       "Final",
		Multiplier: 8,
		OpenDate:   time.Date(2022, 12, 10, 0, 0, 0, 0, time.Local),
		CloseDate:  time.Date(2022, 12, 18, 14, 0, 0, 0, time.Local),
	}
	terceiro := models.Bracket{
		ID:         config.BRKT_TERCERIRO,
		Name:       "Disputa pelo 3º lugar",
		Multiplier: 3,
		OpenDate:   time.Date(2022, 12, 12, 0, 0, 0, 0, time.Local),
		CloseDate:  time.Date(2022, 12, 17, 14, 0, 0, 0, time.Local),
	}
	semi := models.Bracket{
		ID:         config.BRKT_SEMIFINAIS,
		Name:       "Semi Final",
		Multiplier: 3,
		OpenDate:   time.Date(2022, 12, 4, 0, 0, 0, 0, time.Local),
		CloseDate:  time.Date(2022, 12, 14, 18, 0, 0, 0, time.Local),
	}
	quartas := models.Bracket{
		ID:         config.BRKT_QUARTAS,
		Name:       "Quartas",
		Multiplier: 2,
		OpenDate:   time.Date(2022, 12, 2, 0, 0, 0, 0, time.Local),
		CloseDate:  time.Date(2022, 12, 10, 12, 0, 0, 0, time.Local),
	}
	oitavas := models.Bracket{
		ID:         config.BRKT_OITAVAS,
		Name:       "Oitavas",
		Multiplier: 1,
		OpenDate:   time.Date(2022, 11, 30, 12, 0, 0, 0, time.Local),
		CloseDate:  time.Date(2022, 12, 6, 18, 0, 0, 0, time.Local),
	}
	coringa := models.Bracket{
		ID:         config.BRKT_CORINGA,
		Name:       "Rodada Coringa",
		Multiplier: 9,
		OpenDate:   time.Date(2022, 11, 10, 0, 0, 0, 0, time.Local),
		CloseDate:  time.Date(2022, 12, 6, 0, 0, 0, 0, time.Local),
	}

	err := db.Create(&[]models.Bracket{final, terceiro, semi, quartas, oitavas, coringa}).Error

	if err != nil {
		return err
	}

	return nil
}

func nationalTeamsFixtures(db *database.Database) error {
	repo := repositories.NationalTeamRepositoryDb{Db: db}
	bracket := models.Bracket{
		ID: config.BRKT_OITAVAS,
	}
	brazil := models.NationalTeam{
		ID:       config.NT_BRAZIL,
		Name:     "Brazil",
		Brackets: []models.Bracket{bracket},
	}

	_, err := repo.Insert(&brazil)
	if err != nil {
		return err
	}

	france := models.NationalTeam{
		ID:       config.NT_FRANCE,
		Name:     "France",
		Brackets: []models.Bracket{bracket},
	}

	_, err = repo.Insert(&france)
	if err != nil {
		return err
	}
	return nil
}

func matchFixtures(db *database.Database) error {
	matches := []models.Match{
		// Coringa
		{
			ID:        config.MATCH_CORINGA,
			Date:      time.Date(2022, 12, 6, 0, 0, 0, 0, time.Local),
			BracketID: config.BRKT_CORINGA,
		},
		// Oitavas
		{
			ID:        config.MATCH_OITAVAS_01,
			Date:      time.Date(2022, 12, 3, 12, 0, 0, 0, time.Local),
			BracketID: config.BRKT_OITAVAS,
		},
		{
			ID:        config.MATCH_OITAVAS_02,
			Date:      time.Date(2022, 12, 3, 16, 0, 0, 0, time.Local),
			BracketID: config.BRKT_OITAVAS,
		},
		{
			ID:        config.MATCH_OITAVAS_03,
			Date:      time.Date(2022, 12, 4, 12, 0, 0, 0, time.Local),
			BracketID: config.BRKT_OITAVAS,
		},
		{
			ID:        config.MATCH_OITAVAS_04,
			Date:      time.Date(2022, 12, 4, 16, 0, 0, 0, time.Local),
			BracketID: config.BRKT_OITAVAS,
		},
		{
			ID:        config.MATCH_OITAVAS_05,
			Date:      time.Date(2022, 12, 5, 12, 0, 0, 0, time.Local),
			BracketID: config.BRKT_OITAVAS,
		},
		{
			ID:        config.MATCH_OITAVAS_06,
			Date:      time.Date(2022, 12, 5, 16, 0, 0, 0, time.Local),
			BracketID: config.BRKT_OITAVAS,
		},
		{
			ID:        config.MATCH_OITAVAS_07,
			Date:      time.Date(2022, 12, 6, 12, 0, 0, 0, time.Local),
			BracketID: config.BRKT_OITAVAS,
		},
		{
			ID:        config.MATCH_OITAVAS_08,
			Date:      time.Date(2022, 12, 6, 16, 0, 0, 0, time.Local),
			BracketID: config.BRKT_OITAVAS,
		},
		// Quartas
		{
			ID:        config.MATCH_QUARTAS_01,
			Date:      time.Date(2022, 12, 9, 12, 0, 0, 0, time.Local),
			BracketID: config.BRKT_QUARTAS,
		},
		{
			ID:        config.MATCH_QUARTAS_02,
			Date:      time.Date(2022, 12, 9, 16, 0, 0, 0, time.Local),
			BracketID: config.BRKT_QUARTAS,
		},
		{
			ID:        config.MATCH_QUARTAS_03,
			Date:      time.Date(2022, 12, 10, 12, 0, 0, 0, time.Local),
			BracketID: config.BRKT_QUARTAS,
		},
		{
			ID:        config.MATCH_QUARTAS_04,
			Date:      time.Date(2022, 12, 10, 16, 0, 0, 0, time.Local),
			BracketID: config.BRKT_QUARTAS,
		},
		// Semifinais
		{
			ID:        config.MATCH_SEMIFINAIS_01,
			Date:      time.Date(2022, 12, 13, 16, 0, 0, 0, time.Local),
			BracketID: config.BRKT_SEMIFINAIS,
		},
		{
			ID:        config.MATCH_SEMIFINAIS_02,
			Date:      time.Date(2022, 12, 14, 16, 0, 0, 0, time.Local),
			BracketID: config.BRKT_SEMIFINAIS,
		},
		// Terceiro
		{
			ID:        config.MATCH_TERCEIRO,
			Date:      time.Date(2022, 12, 17, 12, 0, 0, 0, time.Local),
			BracketID: config.BRKT_TERCERIRO,
		},
		// Finais
		{
			ID:        config.MATCH_FINAIS,
			Date:      time.Date(2022, 12, 18, 12, 0, 0, 0, time.Local),
			BracketID: config.BRKT_FINAIS,
		},
	}

	err := db.Create(&matches).Error

	if err != nil {
		return err
	}

	return nil
}
