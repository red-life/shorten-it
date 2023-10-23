package services

import (
	"context"
	ports "github.com/red-life/shorten-it/internal/ports/mocks"
	"github.com/red-life/shorten-it/pkg/base62"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKeyGen_GenerateKey(t *testing.T) {
	converter := base62.NewConverter()
	counter := ports.NewCounterRepository()
	kgs := NewKeyGenService(counter, converter)
	ctx := context.Background()
	for i := 1; i <= 10; i++ {
		expectedKey := converter.Encode(int64(i))
		t.Log("key:", expectedKey)
		key, err := kgs.GenerateKey(ctx)
		if err != nil {
			t.Errorf("Expected not to return error but got %s", err)
		}
		assert.Equal(t, int64(i), counter.Count)
		assert.Equalf(t, expectedKey, key, "Expected %s but got %s", expectedKey, key)
	}
}
