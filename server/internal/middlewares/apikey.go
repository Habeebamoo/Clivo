package middlewares

import (
	"log"

	"github.com/Habeebamoo/Clivo/server/internal/config"
	response "github.com/Habeebamoo/Clivo/server/internal/utils"
	"github.com/gin-gonic/gin"
)

func RequireAPIKey() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientApiKey := c.GetHeader("X-API-KEY")
		if clientApiKey == "" {
			response.Abort(c, 401, "API Key Missing", nil)
			return 
		}

		apiKey, err := config.Get("API_KEY")
		if err != nil {
			log.Fatal(err)
		}

		if clientApiKey != apiKey {
			response.Abort(c, 401, "Invalid API Key", nil)
			return 
		}

		//all passed
		c.Next()
	}
}