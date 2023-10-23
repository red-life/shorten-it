package services

import (
	"github.com/red-life/shorten-it/internal/ports"
	"github.com/red-life/shorten-it/pkg/base62"
)

func NewKeyGenService(counter ports.CounterRepository, converter base62.Converter) ports.KeyGenService {
	return &KeyGen{
		counter:   counter,
		converter: converter,
	}
}

type KeyGen struct {
	counter   ports.CounterRepository
	converter base62.Converter
}

func (k KeyGen) GenerateKey() (string, error) {
	//TODO implement me
	panic("implement me")
}
