\c db;

CREATE TABLE users
(
    user_id SERIAL PRIMARY KEY,
    email VARCHAR,
    password VARCHAR NOT NULL,
    profile_picture VARCHAR,
    background_picture VARCHAR,
    account_type_id INTEGER,
    city_id INTEGER,
    verified_user_id INTEGER,
    account_state_id INTEGER,
    auth_type_id INTEGER
);

CREATE TABLE contact
(
    user_id INTEGER REFERENCES users (user_id),
    contact VARCHAR NOT NULL
);

CREATE TABLE user_session
(
    user_session_id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users (user_id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    last_activity TIMESTAMPTZ NOT NULL DEFAULT NOW()
);