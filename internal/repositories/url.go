package repositories

import (
	"context"
	"github.com/red-life/shorten-it/internal/models"
	"github.com/red-life/shorten-it/internal/pkg/customerror"
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

func (U *URL) Save(ctx context.Context, url models.URL) error {
	result := U.db.WithContext(ctx).Create(url)
	return customerror.MapGormToCustomError(result.Error)
}

func (U *URL) GetLongByKey(ctx context.Context, key string) (string, error) {
	url, err := U.get(ctx, models.URL{Key: key})
	return url.Long, err

}

func (U *URL) GetKeyByLong(ctx context.Context, longURL string) (string, error) {
	url, err := U.get(ctx, models.URL{Long: longURL})
	return url.Key, err
}

func (U *URL) get(ctx context.Context, condition models.URL) (models.URL, error) {
	url := new(models.URL)
	result := U.db.WithContext(ctx).Where(condition).First(url)
	return *url, customerror.MapGormToCustomError(result.Error)
}
