CREATE TABLE users (
    id SERIAL PRIMARY KEY, 
    username  TEXT NOT NULL UNIQUE, 
    password TEXT NOT NULL,
    role TEXT NOT NULL DEFAULT 'reguler',
    CONSTRAINT chk_role CHECK (role in ('merchant', 'outlet', 'reguler'))
); 