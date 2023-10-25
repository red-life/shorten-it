package http

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/red-life/shorten-it/internal/pkg/customerror"
	"github.com/red-life/shorten-it/internal/ports"
	"net/http"
)

func RegisterShortenerRoutes(shortener *Shortener, engine *gin.Engine) {
	engine.POST("/shorten", shortener.Shorten)
	engine.GET("/:key", shortener.Redirect)
}

func NewShortenerAdapter(shortenerService ports.ShortenerService, validate *validator.Validate) *Shortener {
	return &Shortener{
		shortenerService: shortenerService,
		validate:         validate,
	}
}

type Shortener struct {
	shortenerService ports.ShortenerService
	validate         *validator.Validate
}

func (s *Shortener) Shorten(c *gin.Context) {
	var request ShortenRequest
	if err := c.BindJSON(&request); err != nil {
		s.returnError(c, customerror.ErrValidation)
		return
	}
	key, err := s.shortenerService.Shorten(c.Request.Context(), request.URL)
	if err != nil {
		s.returnError(c, err)
		return
	}
	c.JSON(http.StatusOK, ShortenResponse{Key: key})
}

func (s *Shortener) Redirect(c *gin.Context) {
	var request RedirectRequest
	if err := c.BindUri(&request); err != nil {
		s.returnError(c, customerror.ErrValidation)
		return
	}
	longURL, err := s.shortenerService.GetLongURL(c.Request.Context(), request.Key)
	if err != nil {
		s.returnError(c, err)
		return
	}
	c.Redirect(http.StatusFound, longURL)
}

func (s *Shortener) returnError(c *gin.Context, err error) {
	c.JSON(
		customerror.MapCustomErrorToHttpStatusCode(err),
		ErrorResponse{err: err},
	)
}
