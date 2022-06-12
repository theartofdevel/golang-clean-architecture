package dto

type BookEvent struct {
	EventType string
	Data      struct {
		Name       string
		Year       int
		AuthorUUID string
		GenreUUID  string
	}
}