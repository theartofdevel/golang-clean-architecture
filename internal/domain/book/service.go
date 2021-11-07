package book

import "context"

type Service interface {
	GetBookByUUID(ctx context.Context, uuid string) *Book
	GetAllBooks(ctx context.Context, limit, offset int) []*Book
	CreateBook(ctx context.Context, dto *CreateBookDTO) *Book
}

type service struct {
	storage Storage
}

func NewService(storage Storage) Service {
	return &service{storage: storage}
}

func (s *service) CreateBook(ctx context.Context, dto *CreateBookDTO) *Book {
	return nil
}

func (s *service) GetBookByUUID(ctx context.Context, uuid string) *Book {
	return s.storage.GetOne(uuid)
}

func (s *service) GetAllBooks(ctx context.Context, limit, offset int) []*Book {
	return s.storage.GetAll(limit, offset)
}