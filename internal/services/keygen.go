package services

import (
	"context"
	"github.com/red-life/shorten-it/internal/ports"
	"sync"
)

func NewKeyGenService(counter ports.CounterRepository, converter ports.Converter) ports.KeyGenService {
	return &KeyGen{
		counter:   counter,
		converter: converter,
		mutex:     sync.Mutex{},
	}
}

type KeyGen struct {
	counter   ports.CounterRepository
	converter ports.Converter
	mutex     sync.Mutex
}

func (k *KeyGen) GenerateKey(ctx context.Context) (string, error) {
	k.mutex.Lock()
	defer k.mutex.Unlock()
	counter, err := k.counter.Increase(ctx)
	if err != nil {
		return "", err
	}
	key := k.converter.Encode(counter)
	return key, nil

}
