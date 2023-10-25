package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/red-life/shorten-it/internal/adapters/http"
	"github.com/red-life/shorten-it/internal/repositories"
	"github.com/red-life/shorten-it/internal/services"
	"github.com/red-life/shorten-it/pkg/base62"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func newRedisClient(host, port, pass string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: pass,
		DB:       0,
	})
}

func newGormDB() *gorm.DB {
	host, port, user, pass, dbName := os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USERNAME"), os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, pass, dbName, port)
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
	serverAddr := fmt.Sprintf("%s:%s", os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT"))
	cacheRDB := newRedisClient(os.Getenv("REDIS_CACHE_HOST"), os.Getenv("REDIS_CACHE_PORT"), os.Getenv("REDIS_CACHE_PASSWORD"))
	counterRDB := newRedisClient(os.Getenv("REDIS_COUNTER_HOST"), os.Getenv("REDIS_COUNTER_PORT"), os.Getenv("REDIS_COUNTER_PASSWORD"))
	db := newGormDB()
	if err != nil {
		panic(err)
	}
	cacheRepo := repositories.NewCache(cacheRDB)
	urlRepo := repositories.NewURLRepository(db)
	counterRepo := repositories.NewCounterRepository(counterRDB)
	urlRepoWithCache := repositories.NewURLWithCacheRepository(cacheRepo, urlRepo)
	converter := base62.NewConverter()
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
