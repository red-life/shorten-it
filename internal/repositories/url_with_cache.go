package repositories

import (
	"context"
	"errors"
	"fmt"
	"github.com/red-life/shorten-it/internal/models"
	"github.com/red-life/shorten-it/internal/pkg/customerror"
	"github.com/red-life/shorten-it/internal/ports"
)

const URLCachePrefix = "URL"

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

func (U *URLWithCache) Save(ctx context.Context, url models.URL) error {
	err := U.urlRepo.Save(ctx, url)
	if err != nil {
		return err
	}
	longKey := fmt.Sprintf("%s:%s:long", URLCachePrefix, url.Key)
	keyKey := fmt.Sprintf("%s:%s:key", URLCachePrefix, url.Long)
	err = U.cache.Set(ctx, longKey, url.Long)
	if err != nil {
		return err
	}
	err = U.cache.Set(ctx, keyKey, url.Key)
	if err != nil {
		return err
	}
	return nil
}

func (U *URLWithCache) GetLongByKey(ctx context.Context, key string) (string, error) {
	keyKey := fmt.Sprintf("%s:%s:long", URLCachePrefix, key)
	long, err := U.cache.Get(ctx, keyKey)
	if err == nil {
		return long, nil
	}
	if !errors.Is(err, customerror.ErrNotFound) {
		return "", err
	}
	long, err = U.urlRepo.GetLongByKey(ctx, key)
	if err != nil {
		return "", err
	}
	err = U.cache.Set(ctx, keyKey, long)
	if err != nil {
		return "", err
	}
	return long, nil
}

func (U *URLWithCache) GetKeyByLong(ctx context.Context, longURL string) (string, error) {
	longKey := fmt.Sprintf("%s:%s:key", URLCachePrefix, longURL)
	key, err := U.cache.Get(ctx, longKey)
	if err == nil {
		return key, nil
	}
	if !errors.Is(err, customerror.ErrNotFound) {
		return "", err
	}
	key, err = U.urlRepo.GetKeyByLong(ctx, longURL)
	if err != nil {
		return "", err
	}
	err = U.cache.Set(ctx, longKey, key)
	if err != nil {
		return "", err
	}
	return key, nil
}
