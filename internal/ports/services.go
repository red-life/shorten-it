package ports

import "context"

type Shortener interface {
	Shorten(ctx context.Context, url string) (string, error)
	GetLongURL(ctx context.Context, key string) (string, error)
}
