package middlewares

import (
	response "github.com/Habeebamoo/Clivo/server/pkg/utils"

	"github.com/gin-gonic/gin"
)

type ReqBody struct {
	Email string `json:"email"`
}

func AuthenticateUser() gin.HandlerFunc {
	return func(c *gin.Context) {

		//real logic goes here


		//fake logic
		var reqBody ReqBody
		if err := c.ShouldBindJSON(&reqBody); err != nil {
			response.Abort(c, 401, "Invalid JSON Format", nil)
			return 
		}

		if reqBody.Email != "habeeb@gmail.com" {
			response.Abort(c, 401, "Invalid Email", nil)
			return 
		}

		c.Set("email", reqBody.Email)
		c.Next()
	}
}

// func VerifyAdmin() gin.HandlerFunc {

// }