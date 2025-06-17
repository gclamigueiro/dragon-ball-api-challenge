CREATE TABLE IF NOT EXISTS characters (
    id INTEGER PRIMARY KEY,
    name VARCHAR NOT NULL, --  It a reserved word, but GORM handles it fine
    ki VARCHAR,
    race VARCHAR
);