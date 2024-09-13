-- +goose Up

CREATE TABLE events (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    userId INTEGER NOT NULL,
    groupId INTEGER NOT NULL,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    time DATETIME NOT NULL,
    FOREIGN KEY (groupId) REFERENCES groups(groupId) ON DELETE CASCADE
);

CREATE TABLE eventsStatus(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    eventId INTEGER NOT NULL,
    userId INTEGER NOT NULL,
    role TEXT,
    pending TEXT CHECK(pending IN ('pending', 'completed', 'rejected')),
    FOREIGN KEY (eventId) REFERENCES events(id) ON DELETE CASCADE
);

-- +goose Down

DROP TABLE events;
DROP TABLE eventsStatus;