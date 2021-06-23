CREATE TABLE merchant_products(
    id SERIAL PRIMARY KEY,
    product_id INT NOT NULL REFERENCES products(id),
    merchant_id INT NOT NULL REFERENCES merchants(id)
);