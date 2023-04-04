package db

import (
	"database/sql"
	"log"
	"path/filepath"
	"runtime"

	"github.com/go-testfixtures/testfixtures/v3"
)

var fixtures *testfixtures.Loader

// SetUpSample 設定假資料
func SetUpSample(db *sql.DB) {
	if fixtures == nil {
		fixtures = initTestFixtures(db)
	}
	if err := fixtures.Load(); err != nil {
		log.Fatal(err)
	}
}

func initTestFixtures(db *sql.DB) *testfixtures.Loader {
	_, file, _, _ := runtime.Caller(0)

	samplePath := filepath.Dir(file) + "/samples"

	fixtures, err := testfixtures.New(
		testfixtures.DangerousSkipTestDatabaseCheck(),
		testfixtures.Database(db),
		testfixtures.Dialect("postgres"),
		testfixtures.UseAlterConstraint(),
		testfixtures.ResetSequencesTo(5000),
		testfixtures.Directory(samplePath),
	)
	if err != nil {
		panic(err)
	}
	return fixtures
}
