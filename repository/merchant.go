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

func (m *merchantRepository) GetMerchantId(userID uint) (*uint, error) {
	var id uint
	err := m.db.Raw("SELECT id from merchants WHERE user_id = ?", userID).Find(&id)
	if err.Error != nil {
		return nil, err.Error
	}
	return &id, nil
}
