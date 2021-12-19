package author

import (
	"ca-library-app/internal/domain/author"
)

type Service interface {
	GetAuthorByUUID(uuid string) *author.Author
	CreateAuthor(dto *author.CreateAuthorDTO) *author.Author
}
