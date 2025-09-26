package handlers

import (
	"github.com/Habeebamoo/Clivo/server/internal/services"
	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	svc services.PostService
}

func NewPostHandler(svc services.PostService) PostHandler {
	return PostHandler{svc: svc}
}

func (pHdl *PostHandler) CreatePost(C *gin.Context) {

}