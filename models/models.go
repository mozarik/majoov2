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
	Outlet   *Outlet   `gorm:"foreignkey:UserID;association_foreignkey:ID"`
	Merchant *Merchant `gorm:"foreignkey:UserID;association_foreignkey:ID"`
}

type Merchant struct {
	gorm.Model
	Product []Product `gorm:"foreignkey:MerchantID;association_foreignkey:ID" json:"product"`
	Outlet  []Outlet  `gorm:"foreignkey:MerchantID;association_foreignkey:ID" json:"outlet"`
	UserID  uint      `json:"user_id"`
}

type Product struct {
	gorm.Model
	Name       string `json:"name"`
	Sku        uint   `json:"sku"`
	Image      string `json:"image"`
	MerchantID uint   `json:"merchant_id"`
}

type ProductOutlet struct {
	gorm.Model
	Product  []Product `gorm:"foreignkey:MerchantID;association_foreignkey:ID"`
	Price    uint
	OutletID uint
}

type Outlet struct {
	gorm.Model
	UserID     uint
	Product    []ProductOutlet `gorm:"foreignkey:OutletID;association_foreignkey:ID"`
	MerchantID uint
}

func Migrate(db *gorm.DB) {
	fmt.Println("Migrating DB")
	db.AutoMigrate(
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
		&User{},
		&Merchant{},
		&Product{},
		&ProductOutlet{},
		&Outlet{},
	)
}
