package usecase_test

import (
	"errors"
	"game-service/internal/entity"
	"game-service/internal/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockRepo struct {
	mockData map[string]*entity.Game
}

func (m *mockRepo) FindAll() ([]entity.Game, error) {
	var games []entity.Game
	for _, v := range m.mockData {
		games = append(games, *v)
	}
	return games, nil
}
func (m *mockRepo) FindByName(name string) (*entity.Game, error) {
	if game, ok := m.mockData[name]; ok {
		return game, nil
	}
	return nil, errors.New("not found")
}
func (m *mockRepo) Create(g *entity.Game) error {
	m.mockData[g.Name] = g
	return nil
}
func (m *mockRepo) Update(name string, g *entity.Game) error {
	m.mockData[name] = g
	return nil
}
func (m *mockRepo) Delete(name string) error {
	delete(m.mockData, name)
	return nil
}

func TestGameUseCase(t *testing.T) {
	mock := &mockRepo{mockData: make(map[string]*entity.Game)}
	uc := usecase.NewGameUseCase(mock)

	// Create
	game := &entity.Game{Name: "Elden Ring", Price: "49.99$", Image: "eldenring.jpg"}
	err := uc.CreateGame(game)
	assert.Nil(t, err)

	// Get
	res, err := uc.GetGame("Elden Ring")
	assert.Nil(t, err)
	assert.Equal(t, "Elden Ring", res.Name)

	// Update
	game.Price = "39.99$"
	uc.UpdateGame("Elden Ring", game)
	assert.Equal(t, "39.99$", mock.mockData["Elden Ring"].Price)

	// Delete
	err = uc.DeleteGame("Elden Ring")
	assert.Nil(t, err)
}
