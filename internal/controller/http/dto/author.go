package dto

type CreateAuthorDTO struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type UpdateAuthorDTO struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
