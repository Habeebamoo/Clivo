package services

import (
	"time"

	"github.com/Habeebamoo/Clivo/server/internal/models"
	"github.com/Habeebamoo/Clivo/server/internal/repositories"
	"github.com/Habeebamoo/Clivo/server/pkg/utils"
)

type ArticleService interface {
	CreateArticle(models.ArticleRequest) (int, error)
}

type ArticleSvc struct {
	repo repositories.ArticleRepository
}

func NewArticleService(repo repositories.ArticleRepository) ArticleService {
	return &ArticleSvc{repo: repo}
}

func (as *ArticleSvc) CreateArticle(articleReq models.ArticleRequest) (int, error) {
	//calculate read time
  readTime := ""

	//upload article image
	articleImage := ""

	//assign article
	article := models.Article{
		ArticleId: utils.GenerateRandomId(),
		AuthorId: articleReq.UserId,
		Title: articleReq.Title,
		Content: articleReq.Content,
		CreatedAt: time.Now(),
		ReadTime: readTime,
		Picture: articleImage,
	}

	return as.repo.CreateArticle(article)

	//notify followers here
}