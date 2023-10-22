package repositories

import (
	"context"
	"errors"
	"fmt"
	"github.com/red-life/shorten-it/internal/pkg/customerror"
	"github.com/red-life/shorten-it/internal/ports"
	"github.com/redis/go-redis/v9"
	"strconv"
)

const CounterPrefix = "COUNTER"

func NewCounterRepository(rdb *redis.Client) ports.CounterRepository {
	return &Counter{
		rdb: rdb,
	}
}

type Counter struct {
	rdb *redis.Client
}

func (c *Counter) GetCounter(ctx context.Context) (int64, error) {
	key := fmt.Sprintf("%s:counter", CounterPrefix)
	val, err := c.rdb.Get(ctx, key).Result()
	if err == nil {
		counter, _ := strconv.ParseInt(val, 10, 64)
		return counter, nil
	}
	err = customerror.MapRedisToCustomError(err)
	if errors.Is(err, customerror.ErrNotFound) {
		return 0, nil
	}
	return 0, err
}

func (c *Counter) Increase(ctx context.Context) (int64, error) {
	key := fmt.Sprintf("%s:counter", CounterPrefix)
	counter, err := c.rdb.Incr(ctx, key).Result()
	return counter, customerror.MapRedisToCustomError(err)
}
