package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mozarik/majoov2/handler"
	"github.com/mozarik/majoov2/middleware"
	model "github.com/mozarik/majoov2/models"
	"gopkg.in/go-playground/validator.v9"
)

func Ping(c echo.Context) error {
	return c.JSON(http.StatusOK, "connected")
}

func main() {
	db, err := model.InitDatabase()
	model.Drop(db)
	if err != nil {
		panic(err)
	}
	model.Migrate(db)

	e := echo.New()

	e.Validator = &middleware.CustomValidator{Validator: validator.New()}
	e.Use(middleware.ContextDB(db))

	e.GET("/", Ping)
	e.POST("/register", handler.RegisterUser)

	e.Logger.Fatal(e.Start(":4001"))
}
