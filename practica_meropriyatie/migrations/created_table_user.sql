DO $$ BEGIN
CREATE TYPE user_role AS ENUM (
        'student',
        'organizer',
        'responsible',
        'admin'
    );
EXCEPTION
    WHEN duplicate_object THEN null;
END $$;

CREATE TABLE IF NOT EXISTS users (
                                     id SERIAL PRIMARY KEY,
                                     name VARCHAR(50) NOT NULL,
    surname VARCHAR(50) NOT NULL,
    login VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(100) NOT NULL,
    phone_number VARCHAR(100) NOT NULL,
    email VARCHAR(150) UNIQUE NOT NULL,
    url_avatar VARCHAR(350),
    user_role user_role NOT NULL DEFAULT 'student',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );