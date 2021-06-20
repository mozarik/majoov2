package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	model "github.com/mozarik/majoov2/models"
	"github.com/mozarik/majoov2/repository"
	"gorm.io/gorm"
)

type RegisterUserBody struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Role     string `json:"role" validate:"required,oneof=merchant outlet"`
}

func RegisterUser(c echo.Context) error {

	// v := validator.New()

	var body RegisterUserBody

	err := c.Bind(&body)
	if err != nil {
		return err
	}

	// Validate struct
	// err = v.Struct(&body)
	err = c.Validate(&body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	db, _ := c.Get("db").(*gorm.DB)

	u := &model.User{
		Username: body.Username,
		Password: body.Password,
		Role:     body.Role,
	}

	repo := repository.NewUserRepository(db)
	err = repo.Register(u)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Success",
		"data":    &body,
	})
}
