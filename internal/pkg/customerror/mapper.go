package customerror

import (
	"errors"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"net/http"
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

func MapCustomErrorToHttpStatusCode(err error) int {
	switch {
	default:
		return http.StatusBadRequest
	}
}
