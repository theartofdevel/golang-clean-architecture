package book

import (
	"ca-library-app/internal/domain/book"
)

type Service interface {
	GetBookByUUID(uuid string) *book.Book
	CreateBook(dto *book.CreateBookDTO) *book.Book
}