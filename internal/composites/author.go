package composites

import (
	"ca-library-app/internal/adapters/api"
	"ca-library-app/internal/adapters/api/author"
	author3 "ca-library-app/internal/adapters/db/author"
	author2 "ca-library-app/internal/domain/author"
)

type AuthorComposite struct {
	Storage author2.Storage
	Service author.Service
	Handler api.Handler
}

func NewAuthorComposite(mongoComposite *MongoDBComposite) (*AuthorComposite, error) {
	storage := author3.NewStorage(mongoComposite.db)
	service := author2.NewService(storage)
	handler := author.NewHandler(service)
	return &AuthorComposite{
		Storage: storage,
		Service: service,
		Handler: handler,
	}, nil
}
