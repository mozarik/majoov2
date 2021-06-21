package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	model "github.com/mozarik/majoov2/models"
	"github.com/mozarik/majoov2/repository"
	"gorm.io/gorm"
)

func GetAllProduct(c echo.Context) {
	panic("Implement me")
}

type RegisterProductBody struct {
	Name  string `json:"name"`
	Sku   uint   `json:"sku"`
	Image string `json:"image"`
}

func NewRegisterProductBody() *RegisterProductBody {
	return &RegisterProductBody{}
}

func RegisterAProduct(c echo.Context) error {
	db, _ := c.Get("db").(*gorm.DB)
	repoProduct := repository.NewProductRepository(db)
	repoUser := repository.NewUserRepository(db)

	cookie, err := c.Cookie("user")
	if err != nil {
		return err
	}

	username := cookie.Value
	_, err = repoUser.GetIDByUsername(username)
	if err != nil {
		return err
	}

	var body RegisterProductBody

	err = c.Bind(&body)
	if err != nil {
		return err
	}

	product := &model.Product{
		Name:  body.Name,
		Sku:   body.Sku,
		Image: body.Image,
	}

	err = repoProduct.CreateProduct(product)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Success",
		"data":    &product,
	})
}
