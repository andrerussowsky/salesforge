-- +goose Up
CREATE TABLE sequences (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    open_tracking_enabled BOOLEAN,
    click_tracking_enabled BOOLEAN
);

CREATE TABLE steps (
    id SERIAL PRIMARY KEY,
    sequence_id INT,
    email_subject VARCHAR(255) NOT NULL,
    email_content TEXT,
    FOREIGN KEY (sequence_id) REFERENCES sequences(id) ON DELETE CASCADE
);
