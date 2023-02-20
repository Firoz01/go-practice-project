package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"database/sql"

	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"

	_ "github.com/uptrace/bun/driver/pgdriver"
)

func getUser(c echo.Context) error {
	id := c.Param("id")

	fmt.Println(id)

	return c.String(http.StatusOK, id)
}

func show(c echo.Context) error {
	team := c.QueryParam("team")
	member := c.QueryParam("member")

	return c.String(http.StatusOK, "team"+team+"member"+member)
}

func databaseConnection() *bun.DB {
	dsn := "postgres://test:abcd1234@localhost:5432/test?sslmode=disable"
	sqldb, err := sql.Open("pg", dsn)
	if err != nil {
		log.Fatal(err)
	}

	db := bun.NewDB(sqldb, pgdialect.New())

	return db
}

func insertDatabase(c echo.Context) error {

	db := databaseConnection()

	user := &User{ID: 2, Name: "Ash"}

	_, err := db.NewInsert().Model(user).Exec(context.Background())

	if err != nil {
		panic(err)
	}

	return c.String(http.StatusAccepted, "Inserted successfully")
}

func selectFromDatabase(c echo.Context) error {
	db := databaseConnection()

	var users []User

	_, err := db.NewSelect().Model(&users).Limit(1).ScanAndCount(context.Background())

	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, users)
}

func main() {

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/users/:id", getUser)

	e.GET("/show", show)

	e.GET("/database", insertDatabase)

	e.GET("/get-data", selectFromDatabase)

	e.Logger.Fatal(e.Start(":1323"))
}
