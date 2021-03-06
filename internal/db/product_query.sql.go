// Code generated by sqlc. DO NOT EDIT.
// source: product_query.sql

package postgres

import (
	"context"
)

const addProduct = `-- name: AddProduct :one
INSERT INTO products (name, image, sku)
VALUES ($1, $2, $3) 
RETURNING id, name, image, sku
`

type AddProductParams struct {
	Name  string `json:"name"`
	Image string `json:"image"`
	Sku   string `json:"sku"`
}

func (q *Queries) AddProduct(ctx context.Context, arg AddProductParams) (Product, error) {
	row := q.queryRow(ctx, q.addProductStmt, addProduct, arg.Name, arg.Image, arg.Sku)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Image,
		&i.Sku,
	)
	return i, err
}

const deleteProduct = `-- name: DeleteProduct :exec
DELETE FROM products
WHERE id = $1
`

func (q *Queries) DeleteProduct(ctx context.Context, id int32) error {
	_, err := q.exec(ctx, q.deleteProductStmt, deleteProduct, id)
	return err
}

const getAllMerchantProducts = `-- name: GetAllMerchantProducts :many
SELECT id, name, image, sku FROM products
WHERE id = ANY(
	SELECT product_id FROM merchant_products
	WHERE merchant_id = $1
)
`

func (q *Queries) GetAllMerchantProducts(ctx context.Context, merchantID int32) ([]Product, error) {
	rows, err := q.query(ctx, q.getAllMerchantProductsStmt, getAllMerchantProducts, merchantID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Product
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Image,
			&i.Sku,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertMerchantProduct = `-- name: InsertMerchantProduct :one
INSERT INTO merchant_products (merchant_id, product_id)
VALUES ($1, $2)
RETURNING merchant_id
`

type InsertMerchantProductParams struct {
	MerchantID int32 `json:"merchantID"`
	ProductID  int32 `json:"productID"`
}

func (q *Queries) InsertMerchantProduct(ctx context.Context, arg InsertMerchantProductParams) (int32, error) {
	row := q.queryRow(ctx, q.insertMerchantProductStmt, insertMerchantProduct, arg.MerchantID, arg.ProductID)
	var merchant_id int32
	err := row.Scan(&merchant_id)
	return merchant_id, err
}
