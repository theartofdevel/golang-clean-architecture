package author

type Storage interface {
	GetOne(uuid string) *Author
	GetAll(limit, offset int) []*Author
	Create(book *Author) *Author
	Delete(book *Author) error
}
