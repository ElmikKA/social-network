-- +goose Up

CREATE TABLE messages (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    userId INTEGER NOT NULL,
    receiverId INTEGER,
    groupId INTEGER,
    message TEXT NOT NULL,
    createdAt DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down

DROP TABLE messages;