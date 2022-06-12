package service

import (
	"context"

	"ca-library-app/internal/domain/entity"
)

type AuthorStorage interface {
	GetOne(id string) entity.Author
	GetAll(ctx context.Context) []entity.Author
	Create(book entity.Author) entity.Author
	Delete(book entity.Author) error
}

type authorService struct {
	storage AuthorStorage
}

func NewAuthorService(storage AuthorStorage) *authorService {
	return &authorService{storage: storage}
}

func (s authorService) Create(ctx context.Context) entity.Author {
	return entity.Author{}
}

func (s authorService) GetByID(ctx context.Context, id string) entity.Author {
	return s.storage.GetOne(id)
}

func (s authorService) GetAll(ctx context.Context) []entity.Author {
	return s.storage.GetAll(ctx)
}
