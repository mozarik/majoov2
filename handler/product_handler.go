package handler

import (
	"encoding/base64"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
	model "github.com/mozarik/majoov2/models"
	"github.com/mozarik/majoov2/repository"
	"gorm.io/gorm"
)

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func GetAllProduct(c echo.Context) {
	panic("Implement me")
}

type RegisterProductBody struct {
	Name  string `json:"name" form:"name"`
	Sku   uint   `json:"sku" form:"sku"`
	Image string `json:"image" form:"image"`
}

func NewRegisterProductBody() *RegisterProductBody {
	return &RegisterProductBody{}
}

func RegisterAProduct(c echo.Context) error {
	db, _ := c.Get("db").(*gorm.DB)

	// THIS IS DRY
	repoProduct := repository.NewProductRepository(db)
	repoUser := repository.NewUserRepository(db)
	repoMerch := repository.NewMerchanRepository(db)

	cookie, err := c.Cookie("user")
	if err != nil {
		return err
	}

	file, err := c.FormFile("image")
	if err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	byteContainer, err := ioutil.ReadAll(src) // why the long names though?
	if err != nil {
		return err
	}

	username := cookie.Value
	userId, err := repoUser.GetIDByUsername(username)
	if err != nil {
		return err
	}

	ImageBase64String := toBase64(byteContainer)
	// log.Println(base64String)

	var body RegisterProductBody

	err = c.Bind(&body)
	if err != nil {
		return err
	}

	merchId, err := repoMerch.GetMerchantId(userId)
	if err != nil {
		return err
	}

	merchProd := []model.MerchantProduct{
		{
			MerchantID: *merchId,
		},
	}

	product := &model.Product{
		Name:            body.Name,
		Sku:             body.Sku,
		Image:           ImageBase64String,
		MerchantProduct: merchProd,
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
