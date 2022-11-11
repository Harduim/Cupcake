package repositories

import (
	"cupcake/app/database"
	"cupcake/app/domain"
	"fmt"
)

type UserPointsRepository interface {
	Insert(userPoints *domain.UserPoints) (*domain.UserPoints, error)
	Find(id string) (*domain.UserPoints, error)
	FindAll() ([]*domain.UserPoints, error)
	Update(userPoints *domain.UserPoints) (*domain.UserPoints, error)
}

type UserPointsRepositoryDb struct {
	Db *database.Database
}

func (repo UserPointsRepositoryDb) Insert(userPoints *domain.UserPoints) (*domain.UserPoints, error) {
	err := repo.Db.Create(userPoints).Error

	if err != nil {
		return nil, err
	}

	return userPoints, nil
}

func (repo UserPointsRepositoryDb) Find(id string) (*domain.UserPoints, error) {
	var userPoints domain.UserPoints

	repo.Db.First(&userPoints, "user_id = ?", id)

	if userPoints.UserID == "" {
		return nil, fmt.Errorf("user does not have points")
	}

	return &userPoints, nil
}

func (repo UserPointsRepositoryDb) FindAll() (*[]domain.UserPoints, error) {
	var userPoints []domain.UserPoints

	repo.Db.Find(&userPoints)

	return &userPoints, nil
}

func (repo UserPointsRepositoryDb) Update(userPoints *domain.UserPoints) (*domain.UserPoints, error) {
	err := repo.Db.Save(&userPoints).Error

	if err != nil {
		return nil, err
	}

	return userPoints, nil
}

func (repo UserPointsRepositoryDb) RunScoreUpdate() error {
	query := "WITH user_normal_score AS (SELECT user_id,\n" +
		"       (CASE WHEN\n" +
		"            b.gol_a = m.gol_a THEN 1 ELSE 0\n" +
		"        END) * multiplier AS goal_a_point,\n" +
		"        (CASE WHEN\n" +
		"            b.gol_b = m.gol_b THEN 1 ELSE 0\n" +
		"        END) * multiplier AS goal_b_point,\n" +
		"        (CASE WHEN\n" +
		"            b.winner_id = m.winner_id THEN 2 ELSE 0\n" +
		"        END) * multiplier AS winner_point,\n" +
		"        (CASE WHEN\n" +
		"            b.gol_a = m.gol_a AND b.gol_b = m.gol_b\n" +
		"            THEN 2 ELSE 0\n" +
		"        END) * multiplier AS all_goal_points,\n" +
		"        (CASE WHEN\n" +
		"            b.winner_id = m.winner_id AND\n" +
		"            b.gol_a = m.gol_a AND\n" +
		"            b.gol_b = m.gol_b\n" +
		"            THEN 3 ELSE 0\n" +
		"        END) * multiplier AS all_goal_and_winner\n" +
		"FROM brackets\n" +
		"INNER JOIN matches m ON brackets.id = m.bracket_id\n" +
		"INNER JOIN bets b ON m.id = b.match_id\n" +
		"),\n" +
		"user_joker_score AS (\n" +
		"    SELECT user_id,\n" +
		"           (CASE WHEN\n" +
		"                (m2.national_team_a_id = j.national_team_a_id\n" +
		"                    AND m2.gol_a = j.gol_a) OR\n" +
		"                (m2.national_team_a_id = j.national_team_b_id\n" +
		"                    AND m2.gol_a = j.gol_b) THEN 21 ELSE 0\n" +
		"             END) AS goal_a_point,\n" +
		"          (CASE WHEN\n" +
		"                (m2.national_team_b_id = j.national_team_b_id\n" +
		"                    AND m2.gol_b = j.gol_b) OR\n" +
		"                (m2.national_team_b_id = j.national_team_a_id\n" +
		"                    AND m2.gol_b = j.gol_a) THEN 21 ELSE 0\n" +
		"             END) AS goal_b_point,\n" +
		"          (CASE WHEN\n" +
		"                m2.winner_id = j.winner_id THEN 42 ELSE 0\n" +
		"             END) AS winner_point\n" +
		"    FROM jokers j\n" +
		"    INNER JOIN brackets b2 ON b2.id = j.bracket_id\n" +
		"    INNER JOIN matches m2 ON b2.id = m2.bracket_id\n" +
		"),\n" +
		"user_groups_score AS (\n" +
		"    SELECT user_id,\n" +
		"           (CASE WHEN\n" +
		"                (m3.national_team_a_id = ug.national_team_id) OR\n" +
		"                (m3.national_team_b_id = ug.national_team_id) THEN 3 ELSE 0\n" +
		"             END) AS team_point\n" +
		"    FROM groups j\n" +
		"    INNER JOIN user_groups ug on j.user_id = ug.groups_user_id\n" +
		"    INNER JOIN brackets b3 on b3.id = j.bracket_id\n" +
		"    INNER JOIN matches m3 on b3.id = m3.bracket_id\n" +
		")\n" +
		"INSERT INTO  user_points (user_id, points) (\n" +
		"    SELECT t.user_id, SUM(score) AS score\n" +
		"    FROM (\n" +
		"        SELECT user_id, SUM(goal_a_point + goal_b_point + winner_point + all_goal_points + all_goal_and_winner) AS score\n" +
		"        FROM user_normal_score\n" +
		"        GROUP BY user_id\n" +
		"            UNION\n" +
		"        SELECT user_id, SUM(goal_a_point + goal_b_point + winner_point) AS score\n" +
		"        FROM user_joker_score\n" +
		"        GROUP BY user_id\n" +
		"            UNION\n" +
		"        SELECT user_id, SUM(u.team_point) AS score\n" +
		"        FROM user_groups_score u\n" +
		"        GROUP BY user_id\n" +
		"    ) t\n" +
		"GROUP BY t.user_id\n" +
		")\n" +
		"ON CONFLICT (user_id) DO UPDATE SET points = EXCLUDED.points;\n" +
		""

	err := repo.Db.Exec(query).Error

	if err != nil {
		return err
	}

	return nil
}
