package ports

import "context"

type ShortenerService interface {
	Shorten(ctx context.Context, url string) (string, error)
	GetLongURL(ctx context.Context, key string) (string, error)
}

type KeyGenService interface {
	GenerateKey(ctx context.Context) (string, error)
}
