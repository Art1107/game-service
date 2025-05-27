package repository

import "game-service/internal/entity"

type GameRepository interface {
	FindAll() ([]entity.Game, error)
	FindByName(name string) (*entity.Game, error)
	Create(game *entity.Game) error
	Update(name string, game *entity.Game) error
	Delete(name string) error
}
