package services

import (
	"context"
	"github.com/red-life/shorten-it/internal/ports"
)

func NewShortenerService(urlRepo ports.URLRepository, kgs ports.KeyGenService) ports.ShortenerService {
	return &Shortener{
		urlRepo: urlRepo,
		kgs:     kgs,
	}
}

type Shortener struct {
	urlRepo ports.URLRepository
	kgs     ports.KeyGenService
}

func (s Shortener) Shorten(ctx context.Context, url string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (s Shortener) GetLongURL(ctx context.Context, key string) (string, error) {
	//TODO implement me
	panic("implement me")
}
