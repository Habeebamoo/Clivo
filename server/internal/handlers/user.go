package handlers

import (
	"github.com/Habeebamoo/Clivo/server/internal/models"
	"github.com/Habeebamoo/Clivo/server/internal/services"
	"github.com/Habeebamoo/Clivo/server/pkg/utils"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	svc services.UserService
}

func NewUserHandler(svc services.UserService) UserHandler {
	return UserHandler{svc}
}

func (usrHdl *UserHandler) GoogleLogin(c *gin.Context) {
	//redirect to google oauth
}

func (usrHdl *UserHandler) GoogleCallBack(c *gin.Context) {


	// sign in user
	user := models.User{}

	token, statusCode, err := usrHdl.svc.SignInUser(user)
	if err != nil {
		utils.Error(c, statusCode, utils.FormatText(err.Error()), nil)
	}

	// send cookie
	utils.SetCookies(c, token)

	utils.Success(c, statusCode, "Signed In Successfully", nil)
}