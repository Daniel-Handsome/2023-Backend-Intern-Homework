package test

import (
	"database/sql"
	"os"
	"testing"
	"time"

	app "github.com/Daniel-Handsome/2023-Backend-intern-Homework"
	"github.com/Daniel-Handsome/2023-Backend-intern-Homework/db"
	"github.com/Daniel-Handsome/2023-Backend-intern-Homework/pb"
	"github.com/Daniel-Handsome/2023-Backend-intern-Homework/test/Setup"
	testDB "github.com/Daniel-Handsome/2023-Backend-intern-Homework/test/db"
	"github.com/Daniel-Handsome/2023-Backend-intern-Homework/utils"
)

var migrDB *sql.DB
var ENV_PATH = "../.env"
var UserClient pb.UserServiceClient
var ArticleClient pb.ArticleServiceClient

func TestMain(t *testing.M) {

	utils.LoadConfig(ENV_PATH)
	//db
	gormDB := db.New()

	// fake data to use
	migrDB, _ = gormDB.DB()
	db.SetUpSample(migrDB)
	testDB.Set(migrDB)

	//clint
	UserClient = Setup.UserClient()
	ArticleClient = Setup.ArticleClient()

	//server
	go app.NewApp(gormDB).Start()

	time.Sleep(2 * time.Second)
	os.Exit(t.Run())
}

func ResetDB() {
	db.SetUpSample(migrDB)
}
