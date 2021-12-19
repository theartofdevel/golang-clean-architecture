package composites

import (
	"ca-library-app/internal/adapters/api/book"
	book3 "ca-library-app/internal/adapters/db/book"
	book2 "ca-library-app/internal/domain/book"
)

type BookComposite struct {
	service book.Service
}

func NewBookComposite() *BookComposite {
	storage := book3.NewStorage()
	return &BookComposite{
		service: book2.NewService(storage),
	}
}
