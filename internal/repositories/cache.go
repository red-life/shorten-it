package repositories

import (
	"context"
	"github.com/red-life/shorten-it/internal/ports"
	"github.com/redis/go-redis/v9"
)

func NewCache(rdb *redis.Client) ports.Cache {
	return &Cache{
		rdb: rdb,
	}
}

type Cache struct {
	rdb *redis.Client
}

func (c Cache) Set(ctx context.Context, key string, value string) error {
	//TODO implement me
	panic("implement me")
}

func (c Cache) Get(ctx context.Context, key string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (c Cache) Delete(ctx context.Context, key string) error {
	//TODO implement me
	panic("implement me")
}
