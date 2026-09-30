package routes

import (
	"github.com/DevLucasHenrique/mymock-backend/internal/handler"
	"github.com/gin-gonic/gin"
)

func InitRoutes(server *gin.Engine, uh *handler.UserHandler) {
	base := server.Group("/api")
	{
		base.GET("/users", uh.ListAllUsers)
	}
}