package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	model "github.com/mozarik/majoov2/models"
	"github.com/mozarik/majoov2/repository"
	"gorm.io/gorm"
)

type RegisterMerchantBopy struct {
	Name string `json:"name"`
}

func RegisterMerchant(c echo.Context) error {
	db, _ := c.Get("db").(*gorm.DB)
	repoUser := repository.NewUserRepository(db)
	repoMerch := repository.NewMerchanRepository(db)

	cookie, err := c.Cookie("user")
	if err != nil {
		return err
	}

	username := cookie.Value

	user_id, err := repoUser.GetIDByUsername(username)
	if err != nil {
		return err
	}

	var body RegisterProductBody

	err = c.Bind(&body)
	if err != nil {
		return err
	}

	m := model.Merchant{
		Name:   body.Name,
		UserID: user_id,
	}

	// TODO RETURN USER ALREADY A MERCHANT
	err = repoMerch.CreateMerchant(&m)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Success",
		"data":    &m,
	})

}
