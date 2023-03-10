package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

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

func insertDatabase(c echo.Context) error {

	db := databaseConnection()

	user := &User{ID: 2, Name: "Ash", Password: "abcd1234"}

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

func employeeInsert(c echo.Context) (err error) {
	db := databaseConnection()

	e := new(Employee)

	if err = c.Bind(e); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	_, err = db.NewInsert().Model(e).Exec(context.Background())

	if err != nil {
		return c.JSON(http.StatusConflict, err.Error())
	}

	return c.JSON(http.StatusAccepted, "insert successfully")

}

func showEmployee(c echo.Context) error {
	db := databaseConnection()

	var employee []Employee

	_, err := db.NewSelect().Model(&employee).ScanAndCount(context.Background())

	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, employee)
}
