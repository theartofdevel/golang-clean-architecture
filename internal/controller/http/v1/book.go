package v1

import (
	"context"
	"encoding/json"
	"net/http"

	"ca-library-app/internal/controller/http/dto"
	"ca-library-app/internal/domain/entity"
	book_usecase "ca-library-app/internal/domain/usecase/book"
	"github.com/julienschmidt/httprouter"
)

const (
	bookURL  = "/books/:book_id"
	booksURL = "/books"
)

type BookUsecase interface {
	CreateBook(ctx context.Context, dto book_usecase.CreateBookDTO) (string, error)
	ListAllBooks(ctx context.Context) []entity.BookView
	GetFullBook(ctx context.Context, id string) entity.FullBook
}

type bookHandler struct {
	bookUsecase BookUsecase
}

func NewBookHandler(bookUsecase BookUsecase) *bookHandler {
	return &bookHandler{bookUsecase: bookUsecase}
}

func (h *bookHandler) Register(router *httprouter.Router) {
	router.GET(booksURL, h.GetAllBooks)
}

func (h *bookHandler) GetAllBooks(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// books := h.bookService.GetAll(context.Background(), 0, 0)
	w.Write([]byte("books"))
	w.WriteHeader(http.StatusOK)
}

func (h *bookHandler) CreateBook(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	var d dto.CreateBookDTO
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		return // error
	}

	// validate

	// MAPPING dto.CreateBookDTO --> book_usecase.CreateBookDTO
	usecaseDTO := book_usecase.CreateBookDTO{
		Name:       "",
		Year:       0,
		AuthorUUID: "",
		GenreUUID:  "",
	}
	book, err := h.bookUsecase.CreateBook(r.Context(), usecaseDTO)
	if err != nil {
		// JSON RPC: TRANSPORT: 200, error: {msg, ..., dev_msg}
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(book))
}
