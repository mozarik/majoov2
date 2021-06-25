package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	echomid "github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"github.com/mozarik/majoov2/auth"
	"github.com/mozarik/majoov2/handler"
	postgres "github.com/mozarik/majoov2/internal/db"
	"github.com/mozarik/majoov2/middleware"
	"gopkg.in/go-playground/validator.v9"
)

func Ping(c echo.Context) error {
	return c.JSON(http.StatusOK, "connected")
}

func InitDatabase() (*postgres.Queries, error) {
	conn, err := sql.Open("postgres", "user=root password=root dbname=root sslmode=disable")
	if err != nil {
		return nil, err
	}

	db := postgres.New(conn)
	fmt.Println("Database Connected")
	return db, nil
}

func main() {

	db, err := InitDatabase()
	if err != nil {
		panic(err)
	}
	e := echo.New()

	e.Validator = &middleware.CustomValidator{Validator: validator.New()}
	e.Use(middleware.ContextDB(db))

	e.GET("/", Ping)
	e.POST("/register", handler.RegisterUser)
	e.GET("/getallusers", handler.GetAllUser)

	// Add this to JWT Group
	// e.POST("/user/updatemerchant", handler.UpdateUserToMerchant)

	e.POST("/login", handler.Login)
	e.POST("/logout", handler.Logout)

	// e.POST("/merchant/register", handler.UpdateUserToMerchant)
	e.POST("/product/add", handler.InsertProduct)
	e.GET("/merchant/products", handler.GetAllMerchantProducts)

	// e.GET("/user", handler.GetCurrentUser)
	adminGroup := e.Group("/admin")

	adminGroup.Use(echomid.JWTWithConfig(echomid.JWTConfig{
		Claims:                  &auth.Claims{},
		SigningKey:              []byte(auth.GetJWTSecret()),
		TokenLookup:             "cookie:access-token",
		ErrorHandlerWithContext: auth.JWTErrorChecker,
	}))

	adminGroup.Use(auth.TokenRefresherMiddleware)
	adminGroup.GET("", handler.Admin)

	merchantGroup := e.Group("/merchant")
	merchantGroup.Use(echomid.JWTWithConfig(echomid.JWTConfig{
		Claims:                  &auth.Claims{},
		SigningKey:              []byte(auth.GetJWTSecret()),
		TokenLookup:             "cookie:access-token",
		ErrorHandlerWithContext: auth.JWTErrorChecker,
	}))

	merchantGroup.POST("/register", handler.UpdateUserToMerchant)

	e.Logger.Fatal(e.Start(":4001"))
}
