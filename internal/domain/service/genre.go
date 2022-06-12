package service

import (
	"context"

	"ca-library-app/internal/domain/entity"
)

type GenreStorage interface {
	GetOne(id string) entity.Genre
	GetAll(limit, offset int) []entity.Genre
	Create(genre entity.Genre) entity.Genre
	Delete(genre entity.Genre) error
}

type genreService struct {
	storage GenreStorage
}

func NewGenreService(storage GenreStorage) *genreService {
	return &genreService{storage: storage}
}

func (s *genreService) Create(ctx context.Context) entity.Genre {
	return entity.Genre{}
}

func (s *genreService) GetByID(ctx context.Context, id string) entity.Genre {
	return s.storage.GetOne(id)
}

func (s *genreService) GetAll(ctx context.Context, limit, offset int) []entity.Genre {
	return s.storage.GetAll(limit, offset)
}
