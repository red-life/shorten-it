package ports

import "context"

func NewCounterRepository() *CounterRepository {
	return &CounterRepository{}
}

type CounterRepository struct {
	Count int64
}

func (c *CounterRepository) Increase(ctx context.Context) (int64, error) {
	c.Count++
	return c.Count, nil
}

func (c *CounterRepository) GetCounter(ctx context.Context) (int64, error) {
	return c.Count, nil
}
