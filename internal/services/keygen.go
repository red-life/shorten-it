package services

import "github.com/red-life/shorten-it/internal/ports"

func NewKeyGenService() ports.KeyGenService {
	return &KeyGen{}
}

type KeyGen struct {
}

func (k KeyGen) GenerateKey() (string, error) {
	//TODO implement me
	panic("implement me")
}
