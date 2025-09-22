package handlers

import (
	"github.com/Habeebamoo/Clivo/server/internal/services"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	svc services.UserService
}

func NewUserHandler(svc services.UserService) UserHandler {
	return UserHandler{svc}
}

func (usrHdl *UserHandler) SignIn(c *gin.Context) {

}