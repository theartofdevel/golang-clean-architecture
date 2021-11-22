package author

import (
	"ca-library-app/internal/domain/author"
)

type authorStorage struct {
}

func NewStorage() author.Storage {
	return &authorStorage{}
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
