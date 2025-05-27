package usecase

import (
	"game-service/internal/entity"
	"game-service/internal/repository"
)

type GameUseCase struct {
	repo repository.GameRepository
}

func NewGameUseCase(r repository.GameRepository) *GameUseCase {
	return &GameUseCase{r}
}

func (u *GameUseCase) GetAllGames() ([]entity.Game, error) {
	return u.repo.FindAll()
}

func (u *GameUseCase) GetGame(name string) (*entity.Game, error) {
	return u.repo.FindByName(name)
}

func (u *GameUseCase) CreateGame(game *entity.Game) error {
	return u.repo.Create(game)
}

func (u *GameUseCase) UpdateGame(name string, game *entity.Game) error {
	return u.repo.Update(name, game)
}

func (u *GameUseCase) DeleteGame(name string) error {
	return u.repo.Delete(name)
}
