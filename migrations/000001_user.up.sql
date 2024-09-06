CREATE  TABLE IF NOT EXISTS users(
    id UUID PRIMARY KEY,
    username VARCHAR UNIQUE,
    email VARCHAR,
    pass_hash VARCHAR(255)
);