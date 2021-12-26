package app

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v8"
	"github.com/yingxv/flashcard-go/src/db"
)

// App
type App struct {
	validate *validator.Validate
	trans    *ut.Translator
	uc       *string
	mongo    *db.MongoClient
	rdb      *redis.Client
}

// New 工厂方法
func New(
	validate *validator.Validate,
	trans *ut.Translator,
	uc *string,
	mongo *db.MongoClient,
	rdb *redis.Client,
) *App {

	return &App{
		validate,
		trans,
		uc,
		mongo,
		rdb,
	}
}
