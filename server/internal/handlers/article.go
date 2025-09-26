package handlers

import (
	"github.com/Habeebamoo/Clivo/server/internal/models"
	"github.com/Habeebamoo/Clivo/server/internal/services"
	"github.com/Habeebamoo/Clivo/server/pkg/utils"
	"github.com/gin-gonic/gin"
)

type ArticleHandler struct {
	svc services.ArticleService
}

func NewArticleHandler(svc services.ArticleService) ArticleHandler {
	return ArticleHandler{svc: svc}
}

func (ah *ArticleHandler) CreateArticle(c *gin.Context) {
	//bind body
	var articleReq models.ArticleRequest
	if err := c.ShouldBindJSON(&articleReq); err != nil {
		utils.Error(c, 400, "Invaild JSON Format", nil)
		return
	}

	//validate request
	if err := articleReq.Validate(); err != nil {
		utils.Error(c, 400, utils.FormatText(err.Error()), nil)
		return
	}

	//call article service
	statusCode, err := ah.svc.CreateArticle(articleReq)
	if err != nil {
		utils.Error(c, statusCode, utils.FormatText(err.Error()), nil)
		return
	}

	utils.Success(c, statusCode, "Article Created Successfully", nil)
}