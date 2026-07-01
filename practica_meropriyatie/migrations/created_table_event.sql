DO $$ BEGIN
CREATE TYPE category_events AS ENUM (
        'КОНФЕРЕНЦИЯ',
        'СУББОТНИК',
        'ОЛИМПИАДА',
        'ПРОФЕССИОНАЛЬНАЯ',
        'ТРЕННИНГ',
        'СПОРТИВНАЯ'
    );
CREATE TYPE event_status AS ENUM (
    'ЗАПЛАНИРОВАНО',
    'ПРОВЕДЕНО'
);
EXCEPTION
    WHEN duplicate_object THEN null;
END $$;


CREATE TABLE IF NOT EXISTS events (
                                     id SERIAL PRIMARY KEY,
                                     name VARCHAR(100),
    description VARCHAR(500),
    date_time TIMESTAMP,
    location VARCHAR(255),
    category_events category_events  DEFAULT 'КОНФЕРЕНЦИЯ',
    status event_status DEFAULT 'ЗАПЛАНИРОВАНО',
    organizer_id INT REFERENCES users(id),
    responsible_id INT REFERENCES users(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );