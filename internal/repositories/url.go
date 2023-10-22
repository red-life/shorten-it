package repositories

import (
	"context"
	"github.com/red-life/shorten-it/internal/models"
	"github.com/red-life/shorten-it/internal/ports"
	"gorm.io/gorm"
)

func NewURLRepository(db *gorm.DB) ports.URLRepository {
	return &URL{
		db: db,
	}
}

type URL struct {
	db *gorm.DB
}

func (U URL) Save(ctx context.Context, url models.URL) error {
	//TODO implement me
	panic("implement me")
}

func (U URL) GetLongByKey(ctx context.Context, key string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (U URL) GetKeyByLong(ctx context.Context, longURL string) (string, error) {
	//TODO implement me
	panic("implement me")
}
