package repositories

import "gorm.io/gorm"

type PostRepository interface {
	CreatePost()
}

type PostRepo struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &PostRepo{db: db}
}

func (pr *PostRepo) CreatePost() {

}