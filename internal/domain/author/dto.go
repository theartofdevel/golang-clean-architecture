package author

type CreateAuthorDTO struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type UpdateBookDTO struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
