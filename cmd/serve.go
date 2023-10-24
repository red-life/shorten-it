package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/red-life/shorten-it/internal/adapters/http"
	"github.com/red-life/shorten-it/internal/repositories"
	"github.com/red-life/shorten-it/internal/services"
	base622 "github.com/red-life/shorten-it/pkg/base62"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func newRedisClient(db int) *redis.Client {
	addr, user, pass := os.Getenv("REDIS_ADDR"), os.Getenv("REDIS_USERNAME"), os.Getenv("REDIS_PASSWORD")
	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Username: user,
		Password: pass,
		DB:       db,
	})
}

func newGormDB() *gorm.DB {
	host, port, user, pass, dbName := os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USERNAME"), os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, port, user, pass, dbName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	serverAddr := os.Getenv("SERVER_ADDR")
	cacheRDB := newRedisClient(0)
	counterRDB := newRedisClient(1)
	db := newGormDB()
	if err != nil {
		panic(err)
	}
	cacheRepo := repositories.NewCache(cacheRDB)
	urlRepo := repositories.NewURLRepository(db)
	counterRepo := repositories.NewCounterRepository(counterRDB)
	urlRepoWithCache := repositories.NewURLWithCacheRepository(cacheRepo, urlRepo)
	converter := base622.NewConverter()
	kgs := services.NewKeyGenService(counterRepo, converter)
	shortenerService := services.NewShortenerService(urlRepoWithCache, kgs)
	validate := validator.New()
	shortenerHTTP := http.NewShortenerAdapter(shortenerService, validate)
	engine := gin.Default()
	http.RegisterShortenerRoutes(shortenerHTTP, engine)
	err = engine.Run(serverAddr)
	if err != nil {
		panic(err)
	}

}
