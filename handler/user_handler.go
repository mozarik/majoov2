package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	model "github.com/mozarik/majoov2/models"
	"github.com/mozarik/majoov2/repository"
	"gorm.io/gorm"
)

// func ReadAllUser()

func GetCurrentUser(c echo.Context) error {
	db, _ := c.Get("db").(*gorm.DB)
	repoUser := repository.NewUserRepository(db)

	cookie, err := c.Cookie("user")
	if err != nil {
		return err
	}

	username := cookie.Value

	u, err := repoUser.ReturnCurrentUser(username)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Something is wrong",
		})
	}

	return c.JSON(http.StatusAccepted, u)

}

func RegisterUser(c echo.Context) error {
	type RegisterUserBody struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
		Role     string `json:"role" validate:"required,oneof=merchant outlet"`
	}

	// v := validator.New()
	db, _ := c.Get("db").(*gorm.DB)
	repo := repository.NewUserRepository(db)
	var body RegisterUserBody

	err := c.Bind(&body)
	if err != nil {
		return err
	}

	// err = v.Struct(&body)
	// Validate body
	status, _ := repo.UsernameIsInDb(body.Username)
	if !status {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Username already exist",
		})
	}

	err = c.Validate(&body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	u := &model.User{
		Username: body.Username,
		Password: body.Password,
		Role:     body.Role,
	}

	err = repo.Register(u)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Success",
		"data":    &body,
	})
}
