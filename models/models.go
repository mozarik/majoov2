package model

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Password string
	Role     string
	Merchant *Merchant `gorm:"foreignkey:UserID;association_foreignkey:ID"`
	Outlet   *Outlet   `gorm:"foreignkey:UserID;association_foreignkey:ID"`
}

type Merchant struct {
	gorm.Model
	UserID          uint              `json:"user_id"`
	Name            string            `json:"name"`
	MerchantProduct []MerchantProduct `gorm:"foreignkey:MerchantID;association_foreignkey:ID" json:"product"`
	Outlet          []Outlet          `gorm:"foreignkey:MerchantID;association_foreignkey:ID" json:"outlet"`
}

type MerchantProduct struct {
	ProductID     uint            `gorm:"primarykey" json:"product_id"`
	MerchantID    uint            `json:"merchant_id"`
	ProductOutlet []ProductOutlet `gorm:"foreignkey:ProductID;association_foreignkey:ProductID"`
}

type Product struct {
	gorm.Model
	Name            string            `json:"name"`
	Sku             uint              `json:"sku"`
	Image           string            `json:"image"`
	MerchantProduct []MerchantProduct `gorm:"foreignkey:ProductID;association_foreignkey:ID"`
}

type ProductOutlet struct {
	gorm.Model
	Outlet    []*Outlet `gorm:"many2many:join_table"`
	ProductID uint
	Outlet_id uint
	Price     uint
}

type Outlet struct {
	gorm.Model
	UserID     uint
	Product    []ProductOutlet `gorm:"many2many:join_table"`
	MerchantID uint
}

func Migrate(db *gorm.DB) {
	fmt.Println("Migrating DB")
	db.AutoMigrate(
		&MerchantProduct{},
		&User{},
		&Merchant{},
		&Product{},
		&ProductOutlet{},
		&Outlet{},
	)
}

func InitDatabase() (*gorm.DB, error) {
	dsn := "host=localhost user=root password=secret dbname=majoo port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	fmt.Println("Database Connected")
	return db, nil
}

func Drop(db *gorm.DB) {
	fmt.Println("Dropping DB")
	db.Migrator().DropTable(
		&MerchantProduct{},
		&User{},
		&Merchant{},
		&Product{},
		&ProductOutlet{},
		&Outlet{},
	)
}
