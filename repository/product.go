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
	panic("not implemented") // TODO: Implement
}

func (p *productRepository) GetAllProduct(username string) error {
	panic("not implemented") // TODO: Implement
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
