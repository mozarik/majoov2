package helper

import (
	"database/sql"
	"fmt"

	postgres "github.com/mozarik/majoov2/internal/db"
)

func InitDatabase() (*postgres.Queries, error) {
	conn, err := sql.Open("postgres", "user=root password=root dbname=root sslmode=disable")
	if err != nil {
		return nil, err
	}

	db := postgres.New(conn)
	fmt.Println("Database Connected")
	return db, nil
}

// func Migrate(db *gorm.DB) {
// 	fmt.Println("Migrating DB")
// 	db.AutoMigrate(
// 		&MerchantProduct{},
// 		&User{},
// 		&Merchant{},
// 		&Product{},
// 		&ProductOutlet{},
// 		&Outlet{},
// 	)
// }

// func Drop(db *gorm.DB) {
// 	fmt.Println("Dropping DB")
// 	db.Migrator().DropTable(
// 		&MerchantProduct{},
// 		&User{},
// 		&Merchant{},
// 		&Product{},
// 		&ProductOutlet{},
// 		&Outlet{},
// 	)
// }
