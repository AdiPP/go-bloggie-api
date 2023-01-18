package infrastructure

import (
	"database/sql"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/sqlite3"
	_ "github.com/mattn/go-sqlite3"
)

func NewQueryBuilder() (database *goqu.Database, err error) {
	var (
		driver = "sqlite3"

		db *sql.DB
	)

	if db, err = sql.Open(driver, "./foo.db"); err != nil {
		return
	}

	goqu.SetDefaultPrepared(true)
	database = goqu.New(driver, db)

	return
}
