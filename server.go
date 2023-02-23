package main

import (
	"net/http"

	"github.com/labstack/echo/v4"

	_ "github.com/uptrace/bun/driver/pgdriver"
)

func main() {

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusAccepted, "Hello, World!")
	})

	e.GET("/users/:id", getUser)

	e.GET("/show", show)

	e.GET("/database", insertDatabase)

	e.GET("/get-data", selectFromDatabase)

	e.POST("/insert-employee", employeeInsert)

	e.GET("/show-employee", showEmployee)

	e.Logger.Fatal(e.Start(":1323"))
}
