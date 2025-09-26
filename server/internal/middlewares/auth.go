package middlewares

import (
	response "github.com/Habeebamoo/Clivo/server/pkg/utils"

	"github.com/gin-gonic/gin"
)

func AuthenticateUser() gin.HandlerFunc {
	return func(c *gin.Context) {

		//real logic goes here


		//fake logic
		email := c.GetHeader("email")

		if email != "habeeb@gmail.com" {
			response.Abort(c, 401, "Invalid Email", nil)
			return 
		}

		c.Set("email", email)
		c.Next()
	}
}

// func VerifyAdmin() gin.HandlerFunc {

// }