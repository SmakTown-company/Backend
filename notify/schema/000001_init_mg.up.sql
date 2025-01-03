CREATE TABLE users (
                id SERIAL PRIMARY KEY,
                email VARCHAR(255) UNIQUE NOT NULL,
                phone VARCHAR(20) UNIQUE NOT NULL,
                user_name varchar(30) UNIQUE  NOT NULL ,
                created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                deleted_at TIMESTAMP DEFAULT NULL,
                verified boolean default false
);
CREATE UNIQUE INDEX idx_users_email ON users (email);

CREATE TABLE push_tokens (
                             id SERIAL PRIMARY KEY,
                             user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
                             subject varchar  not null  default 'У вас новое уведомление',
                             body_text varchar not null default 'Уведомление',
                             created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                             updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                             is_looked boolean default false
);

CREATE INDEX idx_push_tokens_user_id ON push_tokens(user_id);

CREATE TABLE email_tokens(
                             id SERIAL primary key ,
                             email varchar(120) unique REFERENCES  users(email),
                             token varchar(255) unique not null,
                             expires_at timestamp default  current_timestamp+INTERVAL '5 hours',
                             created_at timestamp default current_timestamp
);
CREATE UNIQUE INDEX idx_email_tokens_token ON email_tokens (token);