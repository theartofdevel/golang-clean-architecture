package book

type CreateBookDTO struct {
	Name       string `json:"name"`
	Year       int    `json:"year"`
	AuthorUUID string `json:"author_uuid"`
	GenreUUID string `json:"genre_uuid"`
}

type UpdateBookDTO struct {
	UUID       string `json:"uuid"`
	Name       string `json:"name"`
	Year       int    `json:"year"`
	AuthorUUID string `json:"author_uuid"`
	Busy       bool   `json:"busy"`
	Owner      string `json:"owner"`
}
