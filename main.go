package main

import (
	"log"

	"github.com/Daniel-Handsome/2023-Backend-intern-Homework/db"
	"github.com/Daniel-Handsome/2023-Backend-intern-Homework/utils"
)

var ENV_PATH = ".env"

func main() {
	utils.LoadConfig(ENV_PATH)


	db, err := db.InitDatabase(*utils.App)
	
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	
}


