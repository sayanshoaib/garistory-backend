package config

import (
	"database/sql"
	"fmt"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"os"
)

type DbInstance struct {
	Db *bun.DB
}

var DB DbInstance

func ConnectDb() {
	dsn := fmt.Sprintf(
		"postgres://%s:@localhost:5432/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		//os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	//dsn := "postgres://postgres:@localhost:5432/test?sslmode=disable"
	// dsn := "unix://user:pass@dbname/var/run/postgresql/.s.PGSQL.5432"
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	db := bun.NewDB(sqldb, pgdialect.New())

	DB = DbInstance{
		Db: db,
	}
}
