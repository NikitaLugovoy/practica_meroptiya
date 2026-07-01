CREATE TABLE IF NOT EXISTS events_result (
                                     id SERIAL PRIMARY KEY,
                                     event_id INT REFERENCES events(id),
                                     result VARCHAR(750),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(event_id)
    );