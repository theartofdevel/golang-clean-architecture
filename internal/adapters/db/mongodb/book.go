package mongodb

import (
	"ca-library-app/internal/domain/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type bookStorage struct {
	db *mongo.Database
}

func NewBookStorage(db *mongo.Database) *bookStorage {
	return &bookStorage{db: db}
}

func (bs *bookStorage) GetOne(id string) *entity.Book {
	return nil
}
func (bs *bookStorage) GetAll(limit, offset int) []*entity.Book {
	return nil
}
func (bs *bookStorage) Create(book *entity.Book) *entity.Book {
	return nil
}
func (bs *bookStorage) Delete(book *entity.Book) error {
	return nil
}
