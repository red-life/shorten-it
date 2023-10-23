package services

import (
	"context"
	"github.com/red-life/shorten-it/internal/pkg/customerror"
	ports "github.com/red-life/shorten-it/internal/ports/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestShortener_Shorten(t *testing.T) {
	testCases := []struct {
		name          string
		url           string
		key           string
		isSavedBefore bool
	}{
		{"Normal test case 1", "https://google.com", "google", false},
		{"Normal test case 2", "https://bing.com", "bing", false},
		{"Duplicate test case - must return the saved key", "https://google.com", "google", true},
		{"Duplicate test case 2 - must return the saved key", "https://bing.com", "bing", true},
		{"Domain saved before - must return a new key anyway", "https://google.com/search?q=hi", "google", false},
	}
	urlRepo := ports.NewURLRepository(t)
	kgs := ports.NewKeyGenService(t)
	shortener := NewShortenerService(urlRepo, kgs)
	urlRepo.On("Save", mock.Anything, mock.Anything).Return(nil)
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if !tc.isSavedBefore {
				urlRepo.On("GetKeyByLong", mock.Anything, tc.url).Once().Return("", customerror.ErrNotFound)
				kgs.On("GenerateKey", mock.Anything).Once().Return(tc.key, nil)
			}
			key, err := shortener.Shorten(context.Background(), tc.url)
			assert.Nil(t, err)
			assert.Equal(t, tc.key, key)
			urlRepo.On("GetKeyByLong", mock.Anything, tc.url).Return(key, nil)
		})
	}

}
