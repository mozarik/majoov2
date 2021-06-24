package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	postgres "github.com/mozarik/majoov2/internal/db"
)

func UpdateUserToMerchant(c echo.Context) error {
	db, _ := c.Get("db").(*postgres.Queries)

	username, err := c.Cookie("user")
	if err != nil && username.Value == "" {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": "Need login",
			"error":   err,
		})
	}
	id, err := db.UpdateUserToMerchant(context.Background(), username.Value)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err,
		})
	}

	return c.JSON(http.StatusAccepted, map[string]interface{}{
		"message": fmt.Sprintf("%s with id %d is updated to be a merchant", username.Value, id),
	})
}

// MANUAL TEST AND ITS WORKS
func GetAllUser(c echo.Context) error {
	db, _ := c.Get("db").(*postgres.Queries)

	user, err := db.GetUsers(context.Background())
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err,
		})
	}

	return c.JSON(http.StatusAccepted, user)
}

func IsUserNameExist(username string, db *postgres.Queries) (bool, error) {
	id, err := db.IsUsernameExist(context.Background(), username)
	if id != 0 {
		return true, err
	}
	return false, err
}

type RegisterUserBody struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// MANUAL TEST AND WORK
func RegisterUser(c echo.Context) error {
	db, _ := c.Get("db").(*postgres.Queries)

	var body RegisterUserBody
	err := c.Bind(&body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Bad Request",
		})
	}

	status, _ := IsUserNameExist(body.Username, db)
	if status {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Username already exist",
		})
	}

	err = c.Validate(&body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	user, err := db.CreateUser(context.Background(), postgres.CreateUserParams{
		Username: body.Username,
		Password: body.Password,
	})

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "User is Created",
		"data":    &user,
	})
}
