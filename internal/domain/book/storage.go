package book

type Storage interface {
	GetOne(uuid string) *Book
	GetAll(limit, offset int) []*Book
	Create(book *Book) *Book
	Delete(book *Book) error
}
