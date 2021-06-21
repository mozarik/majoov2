package repository

import (
	model "github.com/mozarik/majoov2/models"
	"gorm.io/gorm"
)

type ProductRepository interface {
	CreateProduct(product *model.Product) error
	GetAllProduct(username string) error
	UpdateProduct()
	DeleteById(id uint) error
	DeleteByIds(id ...uint) error // Batch Delete
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *productRepository {
	return &productRepository{
		db: db,
	}
}

func (p *productRepository) CreateProduct(product *model.Product) error {
	return p.db.Create(product).Save(product).Error
}

func (p *productRepository) GetAllProduct(username string) (*model.Product, error) {
	var product model.Product
	id, err := NewUserRepository(p.db).GetIDByUsername(username)
	if err != nil {
		return nil, err
	}

	err = p.db.Where("merchant_id = ?", id).Find(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *productRepository) UpdateProduct() {
	panic("not implemented") // TODO: Implement
}

func (p *productRepository) DeleteById(id uint) error {
	panic("not implemented") // TODO: Implement
}

func (p *productRepository) DeleteByIds(id ...uint) error {
	panic("not implemented") // TODO: Implement
}
