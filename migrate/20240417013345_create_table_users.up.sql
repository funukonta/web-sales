CREATE TABLE IF NOT EXISTS Users (
    email VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255),
    password VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT into Users (email,name,password) values ('evanroy36@gmail.com','evan roy d','$2a$04$tGoSRhmsLrdPx8QvqViaYOWYuE4G3M2Rj7csqp3XI6p56JI5NYgXS');