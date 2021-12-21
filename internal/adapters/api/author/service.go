package author

import (
	"ca-library-app/internal/domain/author"
)

type Service interface {
	GetAuthorByUUID(uuid string) *author.Author
	CreateAuthor(dto *author.CreateAuthorDTO) *author.Author
<<<<<<< HEAD
}
=======
}
>>>>>>> f504895ef6e7d2a563317d66b70d387239c99c8c
