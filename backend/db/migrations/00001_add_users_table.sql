
-- +goose Up
CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name text NOT NULL UNIQUE,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    firstName TEXT NOT NULL,
    lastName TEXT NOT NULL,
    dateOfBirth DATETIME NOT NULL,
    avatar TEXT DEFAULT '',
    nickname TEXT DEFAULT '',
    aboutMe TEXT DEFAULT '',
    online INTEGER DEFAULT -1,
    privacy TEXT NOT NULL 
);

-- +goose Down
DROP TABLE users;