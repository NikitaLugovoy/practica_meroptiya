CREATE TABLE IF NOT EXISTS groups (
                                     id SERIAL PRIMARY KEY,
                                     name VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );