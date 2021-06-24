package handler

import (
	"context"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"

	"github.com/labstack/echo/v4"
	postgres "github.com/mozarik/majoov2/internal/db"
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

func ReturnImageBase64String(file *multipart.FileHeader) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	byteContainer, err := ioutil.ReadAll(src)
	if err != nil {
		return "", err
	}

	ImageBase64String := toBase64(byteContainer)

	return ImageBase64String, err
}

func InsertProduct(c echo.Context) error {
	db, _ := c.Get("db").(*postgres.Queries)

	file, err := c.FormFile("image")
	if err != nil {
		return err
	}

	Image64String, err := ReturnImageBase64String(file)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   err,
			"message": "Failed to upload file",
		})
	}

	var body postgres.AddProductParams
	err = c.Bind(&body)
	if err != nil {
		return err
	}

	productID, err := db.AddProduct(context.Background(), postgres.AddProductParams{
		Name:  body.Name,
		Image: Image64String,
		Sku:   body.Sku,
	})
	if err != nil {
		return err
	}

	cookie, err := c.Cookie("user")
	if err != nil {
		return err
	}

	var merchantData postgres.InsertMerchantProductParams
	merchantID, err := db.GetMerchantID(context.Background(), cookie.Value)
	if err != nil {
		return err
	}

	merchantData.MerchantID = merchantID
	merchantData.ProductID = productID.ID

	// Insert MerchantProduct (Who have product that we just inserted)
	insertMerchantID, err := db.InsertMerchantProduct(context.Background(), merchantData)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": fmt.Sprintf(("Success inserted %s with %d id Product with id %d"), cookie.Value, insertMerchantID, productID.ID),
	})

}
