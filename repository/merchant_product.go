package repository

import (
	model "github.com/mozarik/majoov2/models"
	"gorm.io/gorm"
)

type merchantProduct struct {
	db *gorm.DB
}

func NewMerchantProduct(db *gorm.DB) *merchantProduct {
	return &merchantProduct{
		db: db,
	}
}

func (m *merchantProduct) MerchantProductFactory(merchantID uint) *model.MerchantProduct {
	var merchantProduct model.MerchantProduct
	merchantProduct.MerchantID = merchantID

	return &merchantProduct
}
