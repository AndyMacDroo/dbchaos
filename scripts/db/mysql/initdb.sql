use db;

CREATE TABLE users
(
    user_id SERIAL PRIMARY KEY,
    email VARCHAR(1),
    password VARCHAR(1) NOT NULL,
    profile_picture VARCHAR(1),
    background_picture VARCHAR(1),
    account_type_id INTEGER,
    city_id INTEGER,
    verified_user_id INTEGER,
    account_state_id INTEGER,
    auth_type_id INTEGER
);

CREATE TABLE contact
(
    user_id INTEGER REFERENCES users (user_id),
    contact VARCHAR(255) NOT NULL
);

CREATE TABLE user_session
(
    user_session_id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users (user_id),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    last_activity TIMESTAMP NOT NULL DEFAULT NOW()
);