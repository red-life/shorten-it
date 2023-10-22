package repositories

import (
	"context"
	"github.com/red-life/shorten-it/internal/ports"
	"github.com/redis/go-redis/v9"
	"time"
)

const DefaultTTL = 5 * time.Minute

func NewCache(rdb *redis.Client) ports.Cache {
	return &Cache{
		rdb: rdb,
	}
}

type Cache struct {
	rdb *redis.Client
}

func (c *Cache) Set(ctx context.Context, key string, value string) error {
	return c.rdb.Set(ctx, key, value, DefaultTTL).Err()
}

func (c *Cache) Get(ctx context.Context, key string) (string, error) {
	pipe := c.rdb.Pipeline()
	pipe.Get(ctx, key)
	pipe.ExpireXX(ctx, key, DefaultTTL)
	cmds, err := pipe.Exec(ctx)
	if err != nil {
		return "", err
	}
	return cmds[0].(*redis.StringCmd).Val(), nil
}

func (c *Cache) Delete(ctx context.Context, key string) error {
	c.rdb.Del(ctx, key) // based on redis docs, if the key doesn't exist it will be ignored, so it won't return any errors (https://redis.io/commands/del/)
	return nil
}
