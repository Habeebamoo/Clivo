package handlers

import (
	"github.com/Habeebamoo/Clivo/server/internal/models"
	"github.com/Habeebamoo/Clivo/server/internal/services"
	"github.com/Habeebamoo/Clivo/server/pkg/utils"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	svc services.AuthService
}

func NewAuthHandler(svc services.AuthService) AuthHandler {
	return AuthHandler{svc}
}

func (ahdl *AuthHandler) GoogleLogin(c *gin.Context) {
	//redirect to google oauth
}

func (ahdl *AuthHandler) GoogleCallBack(c *gin.Context) {


	// sign in user
	user := models.User{}

	token, statusCode, err := ahdl.svc.SignInUser(user)
	if err != nil {
		utils.Error(c, statusCode, utils.FormatText(err.Error()), nil)
	}

	// send cookie
	utils.SetCookies(c, token)

	utils.Success(c, statusCode, "Signed In Successfully", nil)
}