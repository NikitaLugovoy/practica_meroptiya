CREATE TABLE IF NOT EXISTS user_groups (
                                     id SERIAL PRIMARY KEY,
                                     user_id INT REFERENCES users(id) on DELETE CASCADE,
                                     group_id INT REFERENCES groups(id) on DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(user_id,group_id)
    );