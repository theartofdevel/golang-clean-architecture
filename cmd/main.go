package main

import (
	author3 "ca-library-app/internal/adapters/api/author"
	book3 "ca-library-app/internal/adapters/api/book"
	author2 "ca-library-app/internal/adapters/db/author"
	book2 "ca-library-app/internal/adapters/db/book"
	"ca-library-app/internal/domain/author"
	"ca-library-app/internal/domain/book"
)

func main() {
	// entry point
	bookStorage := book2.NewStorage()
	bookService := book.NewService(bookStorage)
	bookHandler := book3.NewHandler(bookService)

	authorStorage := author2.NewStorage()
	authorService := author.NewService(authorStorage)
	auhtorHandler := author3.NewHandler(authorService)


}
