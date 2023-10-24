package http

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/red-life/shorten-it/internal/pkg/customerror"
	ports "github.com/red-life/shorten-it/internal/ports/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestShortener_Shorten(t *testing.T) {
	testCases := []struct {
		name       string
		url        string
		statusCode int
		err        error
		key        string
		call       bool
	}{
		{"normal test", "https://google.com", 200, nil, "google", true},
		{"normal test 2", "https://bing.com", 200, nil, "bing", true},
		{"invalid url - validation error", "pwfwefiweofoew.com/ewqewq3213", http.StatusBadRequest, customerror.ErrValidation, "", false},
	}
	shortenerService := ports.NewShortenerService(t)
	validate := validator.New()
	shortenerAdapter := NewShortenerAdapter(shortenerService, validate)
	engine := gin.New()
	RegisterShortenerRoutes(shortenerAdapter, engine)
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.call {
				shortenerService.EXPECT().Shorten(mock.Anything, tc.url).Once().Return(tc.key, tc.err)
			}
			w := httptest.NewRecorder()
			values := map[string]string{"url": tc.url}
			jsonValue, _ := json.Marshal(values)
			req, _ := http.NewRequest("POST", "/shorten", bytes.NewBuffer(jsonValue))
			req.Header.Set("Content-Type", "application/json")
			engine.ServeHTTP(w, req)
			res := w.Result()
			defer res.Body.Close()
			target := make(map[string]string)
			json.NewDecoder(res.Body).Decode(&target)
			assert.Equalf(t, tc.statusCode, w.Code, "Expected status code %d but got %d", tc.statusCode, w.Code)
			if _, ok := target["err"]; ok {
				assert.Equalf(t, tc.statusCode, w.Code, "Expected error %s but got %s", tc.err, target["err"])
			}
			if _, ok := target["key"]; ok {
				assert.Equalf(t, tc.statusCode, w.Code, "Expected key %s but got %s", tc.key, target["key"])
			}
			shortenerService.AssertExpectations(t)
		})
	}
}
