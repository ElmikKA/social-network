-- +goose Up
CREATE TABLE followers (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    userId INTEGER NOT NULL,
    following INTEGER NOT NULL,
    pending TEXT NOT NULL CHECK(pending IN ('pending', 'completed', 'rejected')),
    FOREIGN KEY (userId) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (following) REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE followers;