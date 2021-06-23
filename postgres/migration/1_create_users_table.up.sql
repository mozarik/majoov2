CREATE TABLE users (
    id SERIAL PRIMARY KEY, 
    username TEXT NOT NULL, 
    password TEXT NOT NULL,
    role TEXT,
    CONSTRAINT chk_role CHECK (role in ('merchant', 'outlet'))
); 