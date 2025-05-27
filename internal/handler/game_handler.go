package handler

import (
	"github.com/gin-gonic/gin"
	"game-service/internal/entity"
	"game-service/internal/usecase"
	"net/http"
)

type GameHandler struct {
	uc *usecase.GameUseCase
}

func NewGameHandler(uc *usecase.GameUseCase) *GameHandler {
	return &GameHandler{uc}
}

func (h *GameHandler) RegisterRoutes(r *gin.Engine) {
	r.GET("/games", h.GetAll)
	r.GET("/games/:name", h.GetOne)
	r.POST("/games", h.Create)
	r.PUT("/games/:name", h.Update)
	r.DELETE("/games/:name", h.Delete)
}

func (h *GameHandler) GetAll(c *gin.Context) {
	games, err := h.uc.GetAllGames()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, games)
}

func (h *GameHandler) GetOne(c *gin.Context) {
	name := c.Param("name")
	game, err := h.uc.GetGame(name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Game not found"})
		return
	}
	c.JSON(http.StatusOK, game)
}

func (h *GameHandler) Create(c *gin.Context) {
	var g entity.Game
	if err := c.ShouldBindJSON(&g); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.uc.CreateGame(&g); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, g)
}

func (h *GameHandler) Update(c *gin.Context) {
	name := c.Param("name")
	var g entity.Game
	if err := c.ShouldBindJSON(&g); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.uc.UpdateGame(name, &g); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, g)
}

func (h *GameHandler) Delete(c *gin.Context) {
	name := c.Param("name")
	if err := h.uc.DeleteGame(name); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
}
