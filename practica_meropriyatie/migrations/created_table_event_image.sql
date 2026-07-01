CREATE TABLE IF NOT EXISTS event_images (
                                     id SERIAL PRIMARY KEY,
                                     event_id INT REFERENCES events(id),
    url_images VARCHAR(350),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );