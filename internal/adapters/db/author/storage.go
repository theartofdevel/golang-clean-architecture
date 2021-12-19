package author

import (
	"ca-library-app/internal/domain/author"
	"go.mongodb.org/mongo-driver/mongo"
)

type authorStorage struct {
	db *mongo.Database
}

func NewStorage(db *mongo.Database) author.Storage {
	return &authorStorage{db: db}
}

func (bs *authorStorage) GetOne(uuid string) *author.Author {
	return nil
}
func (bs *authorStorage) GetAll(limit, offset int) []*author.Author {
	return nil
}
func (bs *authorStorage) Create(book *author.Author) *author.Author {
	return nil
}
func (bs *authorStorage) Delete(book *author.Author) error {
	return nil
}
