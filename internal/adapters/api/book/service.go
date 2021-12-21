package book

import (
	"ca-library-app/internal/domain/book"
)

type Service interface {
	GetBookByUUID(uuid string) *book.Book
	CreateBook(dto *book.CreateBookDTO) *book.Book
<<<<<<< HEAD
}
=======
}
>>>>>>> f504895ef6e7d2a563317d66b70d387239c99c8c
