CREATE TABLE outlet_products(
    id SERIAL PRIMARY KEY,
    product_id INT REFERENCES merchant_products(id),
    outlet_id INT REFERENCES outlets(id)
);