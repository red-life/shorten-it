package ports

import (
	"context"
	"github.com/red-life/shorten-it/internal/models"
)

type URLRepository interface {
	Save(ctx context.Context, url models.URL) error
	GetLongByKey(ctx context.Context, key string) (string, error)
	GetKeyByLong(ctx context.Context, longURL string) (string, error)
}
