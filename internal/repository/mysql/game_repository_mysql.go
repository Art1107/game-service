package mysql

import (
	"database/sql"
	"game-service/internal/entity"
	"game-service/internal/repository"
)

type gameRepo struct {
	db *sql.DB
}

func NewGameRepository(db *sql.DB) repository.GameRepository {
	return &gameRepo{db}
}

func (r *gameRepo) FindAll() ([]entity.Game, error) {
	rows, err := r.db.Query("SELECT name, price, image FROM name")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var games []entity.Game
	for rows.Next() {
		var g entity.Game
		if err := rows.Scan(&g.Name, &g.Price, &g.Image); err != nil {
			return nil, err
		}
		games = append(games, g)
	}
	return games, nil
}

func (r *gameRepo) FindByName(name string) (*entity.Game, error) {
	var g entity.Game
	err := r.db.QueryRow("SELECT name, price, image FROM name WHERE name=?", name).
		Scan(&g.Name, &g.Price, &g.Image)
	if err != nil {
		return nil, err
	}
	return &g, nil
}

func (r *gameRepo) Create(game *entity.Game) error {
	_, err := r.db.Exec("INSERT INTO name (name, price, image) VALUES (?, ?, ?)", game.Name, game.Price, game.Image)
	return err
}

func (r *gameRepo) Update(name string, game *entity.Game) error {
	_, err := r.db.Exec("UPDATE name SET price=?, image=? WHERE name=?", game.Price, game.Image, name)
	return err
}

func (r *gameRepo) Delete(name string) error {
	_, err := r.db.Exec("DELETE FROM name WHERE name=?", name)
	return err
}
