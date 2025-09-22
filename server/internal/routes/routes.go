package routes

import (
	"github.com/Habeebamoo/Clivo/server/internal/handlers"
	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(userHandler handlers.UserHandler) *gin.Engine {
	r := gin.Default()

	return r
}