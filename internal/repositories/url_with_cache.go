package repositories

import (
	"context"
	"github.com/red-life/shorten-it/internal/models"
	"github.com/red-life/shorten-it/internal/ports"
)

func NewURLWithCacheRepository(cache ports.Cache, urlRepo ports.URLRepository) ports.URLRepository {
	return &URLWithCache{
		cache:   cache,
		urlRepo: urlRepo,
	}
}

type URLWithCache struct {
	cache   ports.Cache
	urlRepo ports.URLRepository
}

func (U URLWithCache) Save(ctx context.Context, url models.URL) error {
	//TODO implement me
	panic("implement me")
}

func (U URLWithCache) GetLongByKey(ctx context.Context, key string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (U URLWithCache) GetKeyByLong(ctx context.Context, longURL string) (string, error) {
	//TODO implement me
	panic("implement me")
}
