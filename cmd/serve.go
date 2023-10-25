package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/red-life/shorten-it/internal/adapters/http"
	"github.com/red-life/shorten-it/internal/models"
	"github.com/red-life/shorten-it/internal/repositories"
	"github.com/red-life/shorten-it/internal/services"
	"github.com/red-life/shorten-it/pkg/base62"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func newRedisClient(host, pass string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:6379", host),
		Password: pass,
		DB:       0,
	})
	ping := client.Ping(context.Background())
	if _, err := ping.Result(); err != nil {
		panic(err)
	}
	return client
}

func newGormDB() *gorm.DB {
	user, pass, dbName := os.Getenv("POSTGRES_USERNAME"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB")
	dsn := fmt.Sprintf("host=postgres user=%s password=%s dbname=%s port=5432 sslmode=disable", user, pass, dbName)
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
	cacheRDB := newRedisClient("redis_cache", os.Getenv("REDIS_CACHE_PASSWORD"))
	counterRDB := newRedisClient("redis_counter", os.Getenv("REDIS_COUNTER_PASSWORD"))
	db := newGormDB()
	err = db.AutoMigrate(&models.URL{})
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
	err = engine.Run(":5000")
	if err != nil {
		panic(err)
	}

}
