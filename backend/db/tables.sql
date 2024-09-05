CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name text NOT NULL UNIQUE,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    firstName TEXT NOT NULL,
    lastName TEXT NOT NULL,
    dateOfBirth DATETIME NOT NULL,
    avatar BLOB,
    avatrMimeType TEXT,
    nickname TEXT,
    aboutMe TEXT,
    online INTEGER 
);

INSERT OR IGNORE INTO users (name, email, password, firstname, lastname, dateOfBirth) VALUES ("first","first@gmail.com", "pass","fName", "lName", "2024-09-05 11:11:42") 