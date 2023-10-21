package services

import (
	"context"
	"github.com/red-life/shorten-it/internal/ports"
)

func NewShortener() ports.Shortener {
	return &Shortener{}
}

type Shortener struct {
}

func (s Shortener) Shorten(ctx context.Context, url string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (s Shortener) GetLongURL(ctx context.Context, key string) (string, error) {
	//TODO implement me
	panic("implement me")
}



