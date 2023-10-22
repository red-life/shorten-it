package customerror

import (
	"errors"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func MapGormToCustomError(err error) error {
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return ErrNotFound
	default:
		return err
	}
}

func MapRedisToCustomError(err error) error {
	switch {
	case errors.Is(err, redis.Nil):
		return ErrNotFound
	default:
		return err
	}
}
