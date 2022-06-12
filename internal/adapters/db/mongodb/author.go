package mongodb

import (
	"context"

	"ca-library-app/internal/domain/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type authorStorage struct {
	db *mongo.Database
}

func NewAuthStorage(db *mongo.Database) *authorStorage {
	return &authorStorage{db: db}
}

func (bs *authorStorage) GetOne(id string) *entity.Author {
	return nil
}
func (bs *authorStorage) GetAll(ctx context.Context) []*entity.Author {
	return nil
}
func (bs *authorStorage) Create(book *entity.Author) *entity.Author {
	return nil
}
func (bs *authorStorage) Delete(book *entity.Author) error {
	return nil
}
