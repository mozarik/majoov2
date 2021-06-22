package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	echomid "github.com/labstack/echo/v4/middleware"
	"github.com/mozarik/majoov2/auth"
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

	if err != nil {
		panic(err)
	}

	model.Drop(db)
	model.Migrate(db)

	e := echo.New()

	e.Validator = &middleware.CustomValidator{Validator: validator.New()}
	e.Use(middleware.ContextDB(db))

	e.GET("/", Ping)
	e.POST("/register", handler.RegisterUser)
	e.POST("/login", handler.Login)
	e.POST("/logout", handler.Logout)

	e.POST("/merchant/register", handler.RegisterMerchant)
	e.POST("/product/add", handler.RegisterAProduct)

	e.GET("/user", handler.GetCurrentUser)
	adminGroup := e.Group("/admin")

	adminGroup.Use(echomid.JWTWithConfig(echomid.JWTConfig{
		Claims:                  &auth.Claims{},
		SigningKey:              []byte(auth.GetJWTSecret()),
		TokenLookup:             "cookie:access-token",
		ErrorHandlerWithContext: auth.JWTErrorChecker,
	}))

	adminGroup.Use(auth.TokenRefresherMiddleware)
	adminGroup.GET("", handler.Admin)

	e.Logger.Fatal(e.Start(":4001"))
}
