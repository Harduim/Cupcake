package fixtures

import (
	"cupcake/app/config"
	"cupcake/app/database"
	"cupcake/app/models"
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

	return db.Create(&[]models.Bracket{final, terceiro, semi, quartas, oitavas}).Error

}

func nationalTeamsFixtures(db *database.Database) error {
	teams := []models.NationalTeam{
		{ID: config.NT_ALEMANHA, Name: "Alemanha"},
		{ID: config.NT_ARABIA_SAUDITA, Name: "Arábia_Saudita"},
		{ID: config.NT_ARGENTINA, Name: "Argentina"},
		{ID: config.NT_AUSTRALIA, Name: "Australia"},
		{ID: config.NT_BELGICA, Name: "Bélgica"},
		{ID: config.NT_BRASIL, Name: "Brasil"},
		{ID: config.NT_CAMAROES, Name: "Camarões"},
		{ID: config.NT_CANADA, Name: "Canadá"},
		{ID: config.NT_CATAR, Name: "Catar"},
		{ID: config.NT_COREIA_DO_SUL, Name: "Coreia_do_Sul"},
		{ID: config.NT_COSTA_RICA, Name: "Costa_Rica"},
		{ID: config.NT_CROACIA, Name: "Croácia"},
		{ID: config.NT_DINAMARCA, Name: "Dinamarca"},
		{ID: config.NT_EQUADOR, Name: "Equador"},
		{ID: config.NT_ESPANHA, Name: "Espanha"},
		{ID: config.NT_ESTADOS_UNIDOS, Name: "Estados_Unidos"},
		{ID: config.NT_FRANÇA, Name: "França"},
		{ID: config.NT_GANA, Name: "Gana"},
		{ID: config.NT_HOLANDA, Name: "Holanda"},
		{ID: config.NT_INGLATERRA, Name: "Inglaterra"},
		{ID: config.NT_IRÃ, Name: "Irã"},
		{ID: config.NT_JAPAO, Name: "Japao"},
		{ID: config.NT_MARROCOS, Name: "Marrocos"},
		{ID: config.NT_MEXICO, Name: "México"},
		{ID: config.NT_PAIS_DE_GALES, Name: "País_de_Gales"},
		{ID: config.NT_POLONIA, Name: "Polonia"},
		{ID: config.NT_PORTUGAL, Name: "Portugal"},
		{ID: config.NT_SENEGAL, Name: "Senegal"},
		{ID: config.NT_SERVIA, Name: "Servia"},
		{ID: config.NT_SUIÇA, Name: "Suíça"},
		{ID: config.NT_TUNISIA, Name: "Tunisia"},
		{ID: config.NT_URUGUAI, Name: "Uruguai"},
	}

	return db.Create(&teams).Error

}

func matchFixtures(db *database.Database) error {
	matches := []models.Match{
		// Oitavas
		{
			ID:        config.MATCH_OITAVAS_01,
			Name:      "Vencedor gr. A vs Segundo gr. B",
			Date:      time.Date(2022, 12, 3, 12, 0, 0, 0, time.Local),
			BracketID: config.BRKT_OITAVAS,
		},
		{
			ID:        config.MATCH_OITAVAS_02,
			Name:      "Vencedor gr. C vs Segundo gr. D",
			Date:      time.Date(2022, 12, 3, 16, 0, 0, 0, time.Local),
			BracketID: config.BRKT_OITAVAS,
		},
		{
			ID:        config.MATCH_OITAVAS_03,
			Name:      "Vencedor gr. D vs Segundo gr. C",
			Date:      time.Date(2022, 12, 4, 12, 0, 0, 0, time.Local),
			BracketID: config.BRKT_OITAVAS,
		},
		{
			ID:        config.MATCH_OITAVAS_04,
			Name:      "Vencedor gr. B vs Segundo gr. A",
			Date:      time.Date(2022, 12, 4, 16, 0, 0, 0, time.Local),
			BracketID: config.BRKT_OITAVAS,
		},
		{
			ID:        config.MATCH_OITAVAS_05,
			Name:      "Vencedor gr. E vs Segundo gr. F",
			Date:      time.Date(2022, 12, 5, 12, 0, 0, 0, time.Local),
			BracketID: config.BRKT_OITAVAS,
		},
		{
			ID:        config.MATCH_OITAVAS_06,
			Name:      "Vencedor gr. G vs Segundo gr. H",
			Date:      time.Date(2022, 12, 5, 16, 0, 0, 0, time.Local),
			BracketID: config.BRKT_OITAVAS,
		},
		{
			ID:        config.MATCH_OITAVAS_07,
			Name:      "Vencedor gr. F vs Segundo gr. E",
			Date:      time.Date(2022, 12, 6, 12, 0, 0, 0, time.Local),
			BracketID: config.BRKT_OITAVAS,
		},
		{
			ID:        config.MATCH_OITAVAS_08,
			Name:      "Vencedor gr. H vs Segundo gr. G",
			Date:      time.Date(2022, 12, 6, 16, 0, 0, 0, time.Local),
			BracketID: config.BRKT_OITAVAS,
		},
		// Quartas
		{
			ID:        config.MATCH_QUARTAS_01,
			Name:      "Vencedor OF 5 vs Vencedor OF 6",
			Date:      time.Date(2022, 12, 9, 12, 0, 0, 0, time.Local),
			BracketID: config.BRKT_QUARTAS,
		},
		{
			ID:        config.MATCH_QUARTAS_02,
			Name:      "Vencedor OF 1 vs Vencedor OF 2",
			Date:      time.Date(2022, 12, 9, 16, 0, 0, 0, time.Local),
			BracketID: config.BRKT_QUARTAS,
		},
		{
			ID:        config.MATCH_QUARTAS_03,
			Name:      "Vencedor OF 7 vs Vencedor OF 8",
			Date:      time.Date(2022, 12, 10, 12, 0, 0, 0, time.Local),
			BracketID: config.BRKT_QUARTAS,
		},
		{
			ID:        config.MATCH_QUARTAS_04,
			Name:      "Vencedor OF 3 vs Vencedor OF 4",
			Date:      time.Date(2022, 12, 10, 16, 0, 0, 0, time.Local),
			BracketID: config.BRKT_QUARTAS,
		},
		// Semifinais
		{
			ID:        config.MATCH_SEMIFINAIS_01,
			Name:      "Vencedor QF 2 vs Vencedor QF 1",
			Date:      time.Date(2022, 12, 13, 16, 0, 0, 0, time.Local),
			BracketID: config.BRKT_SEMIFINAIS,
		},
		{
			ID:        config.MATCH_SEMIFINAIS_02,
			Name:      "Vencedor QF 4 vs Vencedor QF 3",
			Date:      time.Date(2022, 12, 14, 16, 0, 0, 0, time.Local),
			BracketID: config.BRKT_SEMIFINAIS,
		},
		// Terceiro
		{
			ID:        config.MATCH_TERCEIRO,
			Name:      "Perdedor SF 1 vs Perdedor SF 2",
			Date:      time.Date(2022, 12, 17, 12, 0, 0, 0, time.Local),
			BracketID: config.BRKT_TERCERIRO,
		},
		// Finais
		{
			ID:        config.MATCH_FINAIS,
			Name:      "Vencedor SF 1 vs Vencedor SF 2",
			Date:      time.Date(2022, 12, 18, 12, 0, 0, 0, time.Local),
			BracketID: config.BRKT_FINAIS,
		},
	}

	return db.Create(&matches).Error

}
