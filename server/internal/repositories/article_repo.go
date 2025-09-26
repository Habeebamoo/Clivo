package repositories

import (
	"fmt"

	"github.com/Habeebamoo/Clivo/server/internal/models"
	"gorm.io/gorm"
)

type ArticleRepository interface {
	CreateArticle(models.Article) (int, error)
}

type ArticleRepo struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) ArticleRepository {
	return &ArticleRepo{db: db}
}

func (ar *ArticleRepo) CreateArticle(article models.Article) (int, error) {
	res := ar.db.Create(&article)
	if res.Error != nil {
		return 500, fmt.Errorf("internal server error")
	}

	return 201, nil
}

func (ar *ArticleRepo) GetArticleById(articleId string) (models.Article, int, error) {
	var article models.Article
	res := ar.db.First(&article, "article_id = ?", articleId)
	if res != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return article, 404, fmt.Errorf("article does not exists")
		}
		return article, 500, fmt.Errorf("internal server error")
	}

	return article, 200, nil
}