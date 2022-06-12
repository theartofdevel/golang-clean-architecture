package book_usecase

import (
	"context"

	"ca-library-app/internal/controller/http/dto"
	"ca-library-app/internal/domain/entity"
)

type Service interface {
	GetAllForList(ctx context.Context) []entity.BookView
	GetByID(ctx context.Context, id string) entity.Book
}

type AuthorService interface {
	GetByID(ctx context.Context, id string) entity.Author
}

type GenreService interface {
	GetByID(ctx context.Context, id string) entity.Genre
}

type bookUsecase struct {
	bookService   Service
	authorService AuthorService
	genreService  GenreService
}

func (u bookUsecase) CreateBook(ctx context.Context, dto dto.CreateBookDTO) (string, error) {
	return "", nil
}

func (u bookUsecase) ListAllBooks(ctx context.Context) []entity.BookView {
	// отобразить список книг с именем Жанра и именем Автора
	return u.bookService.GetAllForList(ctx)
}

func (u bookUsecase) GetFullBook(ctx context.Context, id string) entity.FullBook {
	book := u.bookService.GetByID(ctx, id)
	author := u.authorService.GetByID(ctx, book.AuthorID)
	genre := u.genreService.GetByID(ctx, book.GenreID)

	return entity.FullBook{
		Book:   book,
		Author: author,
		Genre:  genre,
	}
}

// pagination
func (u bookUsecase) GetBooksWithAllAuthors(ctx context.Context, id string) []entity.BookView {
	// Book{Authors: [all authors]}
	// book, author(book_id) -=-
	return nil
}
