package repositories

import (
	"github.com/red-life/shorten-it/internal/ports"
	"github.com/redis/go-redis/v9"
)

func NewCounterRepository(rdb *redis.Client) ports.CounterRepository {
	return &Counter{
		rdb: rdb,
	}
}

type Counter struct {
	rdb *redis.Client
}

func (c Counter) GetCounter() (int, error) {
	//TODO implement me
	panic("implement me")
}

func (c Counter) Increase(i int) error {
	//TODO implement me
	panic("implement me")
}
