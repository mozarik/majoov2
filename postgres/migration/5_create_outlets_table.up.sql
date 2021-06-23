CREATE TABLE outlets(
    id SERIAL NOT NULL PRIMARY KEY,
    merchant_id INT NOT NULL REFERENCES merchants(id),
    user_id INT REFERENCES users(id)
);