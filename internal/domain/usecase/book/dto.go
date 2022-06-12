package book_usecase

type CreateBookDTO struct {
	Name       string
	Year       int
	AuthorUUID string
	GenreUUID  string
}
