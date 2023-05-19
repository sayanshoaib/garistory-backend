package config

import (
	"database/sql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

//type DbInstance struct {
//	Db *bun.DB
//}
//
//var DB DbInstance

func ConnectDb() *bun.DB {
	dsn := "postgres://shoaib:garistory123@db:5432/garistory-db?sslmode=disable"
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	db := bun.NewDB(sqldb, pgdialect.New())

	return db
	//DB = DbInstance{
	//	Db: db,
	//}
}
