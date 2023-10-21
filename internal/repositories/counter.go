package repositories

import "github.com/red-life/shorten-it/internal/ports"

func NewCounterRepository() ports.CounterRepository {
	return &Counter{}
}

type Counter struct {
}

func (c Counter) GetCounter() (int, error) {
	//TODO implement me
	panic("implement me")
}

func (c Counter) Increase(i int) error {
	//TODO implement me
	panic("implement me")
}
