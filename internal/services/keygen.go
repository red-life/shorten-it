package services

import "github.com/red-life/shorten-it/internal/ports"

func NewKeyGenService(counter ports.CounterRepository) ports.KeyGenService {
	return &KeyGen{}
}

type KeyGen struct {
	counter ports.CounterRepository
}

func (k KeyGen) GenerateKey() (string, error) {
	//TODO implement me
	panic("implement me")
}
