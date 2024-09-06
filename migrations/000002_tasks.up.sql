CREATE  TABLE IF NOT EXISTS tasks(
    id UUID  PRIMARY KEY,
    user_id UUID REFERENCES users(id),
    title VARCHAR(255),
    created_at VARCHAR(255)
);