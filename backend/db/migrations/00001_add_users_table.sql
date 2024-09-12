
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
    -- avatarMimeType TEXT DEFAULT '',
    nickname TEXT DEFAULT '',
    aboutMe TEXT DEFAULT '',
    online INTEGER DEFAULT -1
);

-- adds a default row to the users table
INSERT INTO users (name, email, password, firstname, lastname, dateOfBirth) 
VALUES ("first","first@gmail.com", "$2a$10$b/dECtQ9x/udGiu7QKueT.fnck54Ozvnnumh3w67gFcYIW8RqR7d6","fName", "lName", "2024-09-05T11:11:42+00:00");
-- +goose Down
DROP TABLE users;