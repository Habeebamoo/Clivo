package services

import "github.com/Habeebamoo/Clivo/server/internal/repositories"

type PostService interface {
	CreatePost()
}

type PostSvc struct {
	repo repositories.PostRepository
}

func NewPostService(repo repositories.PostRepository) PostService {
	return &PostSvc{repo: repo}
}

func (ps *PostSvc) CreatePost() {

}