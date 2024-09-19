-- +goose Up

CREATE TABLE posts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    userId INTEGER NOT NULL,
    groupId INTEGER NOT NULL,
    creator TEXT NOT NULL,
    title TEXC NOT NULL,
    content TEXT,
    avatar TEXT DEFAULT "",
    createdAt DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    privacy TEXT NOT NULL
);

-- +goose Down
DROP TABLE posts;