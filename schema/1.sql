CREATE DATABASE courses;

CREATE TABLE courses
(
    id INTEGER PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL,
    price_usd INTEGER CHECK(price_usd >= 0),
    price_rub INTEGER CHECK(price_rub >= 0)
);
