-- +goose Up

CREATE TABLE groups (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    userId INTEGER NOT NULL,
    title TEXT NOT NULL UNIQUE,
    description TEXT NOT NULL,
    FOREIGN KEY (userId) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE groupMembers (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    userId INTEGER NOT NULL,
    groupId INTEGER NOT NULL,
    role TEXT,
    pending TEXT CHECK(pending IN ('pending', 'completed','rejected')),
    invitee INTEGER NOT NULL,
    FOREIGN KEY (userId) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (groupId) REFERENCES groups(id) ON DELETE CASCADE,
    CONSTRAINT unique_user_group UNIQUE (userId, groupId)
);

-- +goose Down

DROP TABLE groups;
DROP TABLE groupMembers;