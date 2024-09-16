-- +goose Up

CREATE TABLE notifications (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    userId INTEGER NOT NULL,
    content TEXT NOT NULL,
    type TEXT NOT NULL,
    idRef INTEGER NOT NULL,
    createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (userId) REFERENCES users(id) ON DELETE CASCADE 
);

-- +goose Down
DROP TABLE notifications;

