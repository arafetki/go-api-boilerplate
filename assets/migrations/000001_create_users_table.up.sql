CREATE TABLE
    IF NOT EXISTS users (
        id serial PRIMARY KEY,
        email varchar(255) UNIQUE NOT NULL,
        name varchar(255) NOT NULL,
        is_verified boolean NOT NULL DEFAULT false,
        created_at timestamptz NOT NULL DEFAULT now (),
        updated_at timestamptz
    )