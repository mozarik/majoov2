-- name: GetAllMerchantProducts :many
SELECT * FROM products
WHERE id = ANY(
	SELECT product_id FROM merchant_products
	WHERE merchant_id = $1
);

-- name: AddProduct :one
INSERT INTO products (name, image, sku)
VALUES ($1, $2, $3) 
RETURNING *;

-- name: DeleteProduct :exec
DELETE FROM products
WHERE id = $1;

-- name: InsertMerchantProduct :one
INSERT INTO merchant_products (merchant_id, product_id)
VALUES ($1, $2)
RETURNING merchant_id;

