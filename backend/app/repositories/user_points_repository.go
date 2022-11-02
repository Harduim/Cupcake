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
	query := "WITH user_score AS " +
		"(SELECT user_id," +
		"	(CASE WHEN b.gol_a = m.gol_a THEN 1 ELSE 0 END) * multiplier AS goal_a_point, \n" +
		"	(CASE WHEN b.gol_b = m.gol_b THEN 1 ELSE 0 END) * multiplier AS goal_b_point, \n" +
		"	(CASE WHEN b.winner_id = m.winner_id THEN 2 ELSE 0 END) * multiplier AS winner_point, \n" +
		"	(CASE WHEN b.gol_a = m.gol_a AND b.gol_b = m.gol_b THEN 2 ELSE 0 END) * multiplier AS all_goal_points, \n" +
		"	(CASE WHEN b.winner_id = m.winner_id " +
		"		AND b.gol_a = m.gol_a " +
		"		AND b.gol_b = m.gol_b " +
		"	THEN 3 ELSE 0 END) * multiplier AS all_goal_and_winner \n" +
		"FROM brackets \n" +
		"INNER JOIN matches m on brackets.id = m.bracket_id \n" +
		"INNER JOIN bets b on m.id = b.match_id) \n" +
		"INSERT INTO user_points (user_id, points) \n" +
		"(SELECT user_id, SUM(goal_a_point + goal_b_point + winner_point + all_goal_points + all_goal_and_winner) AS score \n" +
		"FROM user_score \n" +
		"GROUP BY user_id) \n" +
		"ON CONFLICT (user_id) DO UPDATE SET points = EXCLUDED.points;"

	err := repo.Db.Exec(query).Error

	if err != nil {
		return err
	}

	return nil
}
