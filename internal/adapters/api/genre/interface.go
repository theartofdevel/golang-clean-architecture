package genre

import (
	"ca-library-app/internal/domain/genre"
	"context"
)

type Service interface {
	GetByUUID(ctx context.Context, uuid string) *genre.Genre
	GetAll(ctx context.Context, limit, offset int) []*genre.Genre
	Create(ctx context.Context, dto *CreateGenreDTO) *genre.Genre
}
