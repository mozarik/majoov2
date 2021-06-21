package repository

import (
	model "github.com/mozarik/majoov2/models"
	"gorm.io/gorm"
)

type merchantRepository struct {
	db *gorm.DB
}

func NewMerchanRepository(db *gorm.DB) *merchantRepository {
	return &merchantRepository{
		db: db,
	}
}

func (m *merchantRepository) CreateMerchant(merchant *model.Merchant) error {
	return m.db.Create(merchant).Save(merchant).Error
}
