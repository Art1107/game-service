package main

import (
	"github.com/gin-gonic/gin"
	"game-service/db"
	"game-service/internal/handler"
	"game-service/internal/repository/mysql"
	"game-service/internal/usecase"
)

func main() {
	dbConn := db.GetDB()
	repo := mysql.NewGameRepository(dbConn)
	uc := usecase.NewGameUseCase(repo)
	h := handler.NewGameHandler(uc)

	r := gin.Default()
	h.RegisterRoutes(r)
	r.Run(":8080")
}
