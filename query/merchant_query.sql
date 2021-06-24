-- name: GetMerchantID :one
SELECT id FROM merchants
WHERE user_id = $1;