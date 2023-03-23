package main

import (
	"fmt"

	"github.com/Daniel-Handsome/2023-Backend-intern-Homework/database/driver/postgres"
	"github.com/Daniel-Handsome/2023-Backend-intern-Homework/utils"
)

var ENV_PATH = ".env"

func main() {
	utils.LoadConfig(ENV_PATH)

	gorm := postgres.New()

	
}
