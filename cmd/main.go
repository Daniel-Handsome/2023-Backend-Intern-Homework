package main

import (
	app "github.com/Daniel-Handsome/2023-Backend-intern-Homework"
	"github.com/Daniel-Handsome/2023-Backend-intern-Homework/db"
	"github.com/Daniel-Handsome/2023-Backend-intern-Homework/utils"
)

var ENV_PATH = ".env"

func main() {
	utils.LoadConfig(ENV_PATH)

	// cache
	//redis := cache.NewRedisStore(cache.New())

	//db
	gorm := db.New()

	// fake data to use
	sqlDB, _ := gorm.DB()
	db.SetUpSample(sqlDB)

	app.NewApp(gorm).Start()
}
