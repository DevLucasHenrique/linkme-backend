package main

import (
	"log"
	"os"

	"github.com/DevLucasHenrique/mymock-backend/internal/handler"
	"github.com/DevLucasHenrique/mymock-backend/internal/infrastructure/database"
	"github.com/DevLucasHenrique/mymock-backend/internal/repository"
	"github.com/DevLucasHenrique/mymock-backend/internal/routes"
	"github.com/DevLucasHenrique/mymock-backend/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	server := gin.Default()


	db, err := database.NewPostgresConnection(os.Getenv("DATABASE_URI"))

	if err!=nil {
		log.Fatalf("Can't connect with database")
	}

	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUsecase)


	routes.InitRoutes(
		server,
		userHandler,
	)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8000"
	}

	server.Run(":" + PORT)
}
