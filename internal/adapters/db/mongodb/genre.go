package mongodb

import (
	"context"

	"ca-library-app/internal/domain/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type genreStorage struct {
	db *mongo.Database
}

func NewGenreStorage(db *mongo.Database) *genreStorage {
	return &genreStorage{db: db}
}

func (bs *genreStorage) GetOne(id string) *entity.Genre {
	return nil
}
func (bs *genreStorage) GetAll(ctx context.Context) []*entity.Genre {
	return nil
}
func (bs *genreStorage) Create(book *entity.Genre) *entity.Genre {
	return nil
}
func (bs *genreStorage) Delete(book *entity.Genre) error {
	return nil
}
