package v1

import (
	"ca-library-app/internal/___backup/adapters/api/author"
	"ca-library-app/internal/backup/adapters/api"
	"github.com/julienschmidt/httprouter"
)

type handler struct {
	authorService author.Service
}

func NewHandler(service author.Service) api.Handler {
	return &handler{authorService: service}
}

func (h *handler) Register(router *httprouter.Router) {
	panic("implement me")
}
