
CREATE TRIGGER delete_on_rejected
AFTER UPDATE ON followers
FOR EACH ROW
WHEN NEW.pending = 'rejected'
BEGIN
    DELETE FROM followers WHERE id = NEW.id;
END;


CREATE TRIGGER delete_expired_sessions AFTER INSERT ON sessions FOR EACH ROW BEGIN
DELETE FROM sessions
WHERE
    expiresAt < CURRENT_TIMESTAMP;
END;


CREATE TRIGGER delete_notification_on_update
AFTER UPDATE OF pending ON followers
FOR EACH ROW
WHEN NEW.pending IN ('completed', 'rejected')
BEGIN
    DELETE FROM notifications WHERE idRef = NEW.id;
END;

CREATE TRIGGER delete_rejected_group_members
AFTER UPDATE ON groupMembers
FOR EACH ROW
WHEN NEW.pending = 'rejected'
BEGIN
    DELETE FROM groupMembers WHERE id = NEW.id;
END;

CREATE TRIGGER set_groupMembers_role_and_pending
AFTER INSERT ON groupMembers
FOR EACH ROW
BEGIN
    UPDATE groupMembers
    SET role = CASE
        WHEN (SELECT userId FROM groups WHERE id = NEW.groupId) = NEW.userId
        THEN 'owner'
        ELSE 'member'
    END,
    pending = CASE
        WHEN (SELECT userId FROM groups WHERE id = NEW.groupId) = NEW.userId
        THEN 'completed'
        ELSE 'pending'
    END
    WHERE id = NEW.id;
END;


CREATE TRIGGER set_event_role_and_pending
AFTER INSERT ON eventsStatus
FOR EACH ROW
BEGIN
    -- Update the inserted row with the correct role and pending values
    UPDATE eventsStatus
    SET role = CASE
        WHEN (SELECT userId FROM events WHERE id = NEW.eventId) = NEW.userId
        THEN 'owner'
        ELSE 'member'
    END,
    pending = CASE
        WHEN (SELECT userId FROM events WHERE id = NEW.eventId) = NEW.userId
        THEN 'completed'
        ELSE 'pending'
    END
    WHERE id = NEW.id;
END;


CREATE TRIGGER delete_rejected_event_status
AFTER UPDATE ON eventsStatus
FOR EACH ROW
WHEN NEW.pending = 'rejected'
BEGIN
    DELETE FROM eventsStatus WHERE id = NEW.id;
END;
