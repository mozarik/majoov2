package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mozarik/majoov2/auth"
	postgres "github.com/mozarik/majoov2/internal/db"
)

func Logout(c echo.Context) error {
	c.SetCookie(&http.Cookie{
		Name:     "access-token",
		Value:    "",
		HttpOnly: true,
	})
	c.SetCookie(&http.Cookie{
		Name:     "refresh-token",
		Value:    "",
		HttpOnly: true,
	})
	c.SetCookie(&http.Cookie{
		Name:     "user",
		Value:    "",
		HttpOnly: true,
	})

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Logout success",
	})
}

// WORK. MANUAL TESTED
func Login(c echo.Context) error {
	db, _ := c.Get("db").(*postgres.Queries)

	var body auth.RegisterLoginBody
	// Parse the submitted data and fill the User struct with the data from the SignIn form.
	if err := c.Bind(&body); err != nil {
		return err
	}

	status, _ := IsUserNameExist(body.Username, db)
	if !status {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Username doesn't exists",
		})
	}

	userPassword, err := db.GetUserPassword(context.Background(), body.Username)
	if err != nil {
		return err
	}

	if userPassword != body.Password {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "Wrong Password",
		})
	}

	err = auth.GenerateTokensAndSetCookies(body, c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Token is incorrect")
	}

	return c.Redirect(http.StatusMovedPermanently, "/admin")
}

// func Login(c echo.Context) error {

// 	db, _ := c.Get("db").(*gorm.DB)
// 	repo := repository.NewUserRepository(db)
// 	var body auth.RegisterLoginBody
// 	// Parse the submitted data and fill the User struct with the data from the SignIn form.
// 	if err := c.Bind(&body); err != nil {
// 		return err
// 	}

// 	// Validate is Username is in DB or not
// 	status, _ := repo.UsernameIsInDb(body.Username)
// 	if status {
// 		return c.JSON(http.StatusUnauthorized, map[string]string{
// 			"message": "Username doesn't exists",
// 		})
// 		// return echo.NewHTTPError(http.StatusUnauthorized, "Password is incorrect")
// 	}

// 	userPassword, err := repo.GetPassword(body.Username)
// 	if err != nil {
// 		return err
// 	}

// 	if userPassword != body.Password {
// 		return c.JSON(http.StatusUnauthorized, map[string]string{
// 			"message": "Wrong Password",
// 		})
// 	}

// 	err = auth.GenerateTokensAndSetCookies(body, c)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusUnauthorized, "Token is incorrect")
// 	}

// 	return c.Redirect(http.StatusMovedPermanently, "/admin")
// }

func Admin(c echo.Context) error {
	// Gets user cookie.
	userCookie, _ := c.Cookie("user")
	return c.String(http.StatusOK, fmt.Sprintf("Hi, %s! You have been authenticated!", userCookie.Value))
}
