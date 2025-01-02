CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       email VARCHAR(255) UNIQUE NOT NULL,
                       phone VARCHAR(20) UNIQUE NOT NULL,
                       user_name varchar(30) UNIQUE  NOT NULL ,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       deleted_at TIMESTAMP DEFAULT NULL

);
CREATE TABLE push_tokens (
                             id SERIAL PRIMARY KEY,
                             user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
                             token VARCHAR(255) NOT NULL,
                             device_type VARCHAR(50) NOT NULL,
                             created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                             updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_push_tokens_user_id ON push_tokens(user_id);