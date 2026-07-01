DO $$ BEGIN
CREATE TYPE participants_status AS ENUM (
        'ПРИШЁЛ',
        'ОТСУТСТВОВАЛ'
    );
EXCEPTION
    WHEN duplicate_object THEN null;
END $$;

CREATE TABLE IF NOT EXISTS event_participants (
                                    id SERIAL PRIMARY KEY,

                                    event_id INT REFERENCES events(id) ON DELETE CASCADE,
                                    user_id INT REFERENCES users(id) ON DELETE CASCADE,
                                    participants_status participants_status DEFAULT 'ПРИШЁЛ',
                                    registered_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

                                    UNIQUE(event_id, user_id)
);