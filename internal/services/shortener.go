package services

import (
	"context"
	"errors"
	"github.com/red-life/shorten-it/internal/models"
	"github.com/red-life/shorten-it/internal/pkg/customerror"
	"github.com/red-life/shorten-it/internal/ports"
	"golang.org/x/sync/singleflight"
)

func NewShortenerService(urlRepo ports.URLRepository, kgs ports.KeyGenService) ports.ShortenerService {
	return &Shortener{
		urlRepo: urlRepo,
		kgs:     kgs,
		group:   singleflight.Group{},
	}
}

type Shortener struct {
	urlRepo ports.URLRepository
	kgs     ports.KeyGenService
	group   singleflight.Group
}

func (s *Shortener) Shorten(ctx context.Context, url string) (string, error) {
	key, err := s.urlRepo.GetKeyByLong(ctx, url)
	if err != nil && errors.Is(err, customerror.ErrNotFound) {
		return s.generateAndSave(ctx, url)
	}
	if err == nil && key != "" {
		return key, nil
	}
	return "", err
}

func (s *Shortener) GetLongURL(ctx context.Context, key string) (string, error) {
	v, err, _ := s.group.Do(key, func() (interface{}, error) {
		return s.urlRepo.GetLongByKey(ctx, key)
	})
	return v.(string), err
}

func (s *Shortener) generateAndSave(ctx context.Context, url string) (string, error) {
	key, err := s.kgs.GenerateKey(ctx)
	if err != nil {
		return "", err
	}
	urlModel := models.URL{
		Long: url,
		Key:  key,
	}
	// TODO: implement the retry pattern
	err = s.urlRepo.Save(ctx, urlModel)
	if err != nil {
		return "", err
	}
	return key, nil
}
