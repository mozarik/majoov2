CREATE TABLE merchants (
    id SERIAL PRIMARY KEY, 
    name TEXT NOT NULL,
    user_id SERIAL NOT NULL REFERENCES users(id)
);