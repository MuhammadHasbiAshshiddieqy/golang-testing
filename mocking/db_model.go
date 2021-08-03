package mocking

import (
	"errors"
	"math/rand"
)


type Post struct {
	ID    int64  `json:"id"`
	Title string `json:"Title"`
	Text  string `json:"Text"`
}

type PostRepository interface {
	Save(post *Post) (*Post, error)
	FindAll() ([]Post, error)
}

type PostService interface {
	Validate(post *Post) error
	Create(post *Post) (*Post, error)
	FindAll() ([]Post, error)
}

type service struct{}

var (
	repo PostRepository
)

func NewPostService(repository PostRepository) PostService {
	repo = repository
	return &service{}
}

func (*service) Validate(post *Post) error {
	if post == nil {
		err := errors.New("The post is empty")
		return err
	}
	if post.Title == "" {
		err := errors.New("The post title is empty")
		return err
	}
	return nil
}

func (*service) Create(post *Post) (*Post, error) {
	post.ID = rand.Int63()
	return repo.Save(post)
}

func (*service) FindAll() ([]Post, error) {
	return repo.FindAll()
}
