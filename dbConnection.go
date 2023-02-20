package main

import (
	"log"

	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"

	_ "github.com/uptrace/bun/driver/pgdriver"
)

func databaseConnection() *bun.DB {
	dsn := "postgres://test:abcd1234@localhost:5432/test?sslmode=disable"
	sqldb, err := sql.Open("pg", dsn)
	if err != nil {
		log.Fatal(err)
	}

	db := bun.NewDB(sqldb, pgdialect.New())

	return db
}
