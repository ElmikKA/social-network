CREATE TRIGGER delete_expired_sessions AFTER INSERT ON sessions FOR EACH ROW BEGIN
DELETE FROM sessions
WHERE
    expiresAt < CURRENT_TIMESTAMP;
END;


CREATE TRIGGER delete_on_rejected
AFTER UPDATE ON followers
FOR EACH ROW
WHEN NEW.pending = 'rejected'
BEGIN
    DELETE FROM followers WHERE id = NEW.id;
END;

CREATE TRIGGER send_follow_notification
AFTER INSERT ON followers
FOR EACH ROW
BEGIN
    INSERT INTO notifications (userId, content, type, idRef)
    SELECT
        NEW.following,
        (SELECT name FROM users WHERE id = NEW.userId) || ' has sent you a follow request',
        'f_ref',
        NEW.id
    WHERE NEW.pending = 'pending';
END;

CREATE TRIGGER delete_notification_on_update
AFTER UPDATE OF pending ON followers
FOR EACH ROW
WHEN NEW.pending IN ('completed', 'rejected')
BEGIN
    DELETE FROM notifications WHERE idRef = NEW.id;
END;




-- groupMembers triggers

CREATE TRIGGER delete_rejected_group_members
AFTER UPDATE ON groupMembers
FOR EACH ROW
WHEN NEW.pending = 'rejected'
BEGIN
    DELETE FROM groupMembers WHERE id = NEW.id;
END;


CREATE TRIGGER add_completed_users_to_events_as_pending
AFTER UPDATE ON groupMembers
FOR EACH ROW
WHEN NEW.pending = 'completed'
BEGIN
    INSERT INTO eventsStatus (eventId, userId, role, pending)
    SELECT 
        e.id,                
        NEW.userId,          
        'member',            
        'pending'            
    FROM events e
    WHERE e.groupId = NEW.groupId; 
END;


CREATE TRIGGER set_groupMembers_role_and_pending_and_create_notification
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

    INSERT INTO notifications (userId, content, type, idRef)
    SELECT 
        (SELECT userId FROM groups WHERE id = NEW.groupId), 
        (SELECT name FROM users WHERE id = NEW.userId) || ' has sent a request to join your group ' || 
        (SELECT title FROM groups WHERE id = NEW.groupId),  
        'g_ref', 
        NEW.id  
    WHERE NEW.invitee = 0;

    INSERT INTO notifications (userId, content, type, idRef)
    SELECT 
        NEW.userId,
        'You have been added to the group ' || (SELECT title FROM groups WHERE id = NEW.groupId),
        'gi_ref',
        NEW.id
    WHERE NEW.invitee = 1;
END;




CREATE TRIGGER delete_notification_on_groupMemberUpdate
AFTER UPDATE OF pending ON groupMembers
FOR EACH ROW
WHEN NEW.pending IN ('completed', 'rejected')
BEGIN
    DELETE FROM notifications WHERE idRef = NEW.id;
END;



-- triggers for events and eventsStatus

CREATE TRIGGER set_group_member_statuses
AFTER INSERT ON events
FOR EACH ROW
BEGIN
    INSERT INTO eventsStatus (eventId, userId, role, pending)
    VALUES( 
        NEW.id,          
        NEW.userId,    
        'owner',        
        'completed'    
    );
    
    INSERT INTO eventsStatus (eventId, userId, role, pending)
    SELECT 
        NEW.id,                 
        gm.userId,             
        'member',              
        'pending'              
    FROM groupMembers gm
    WHERE gm.groupId = NEW.groupId 
    AND gm.pending = 'completed'  
    AND gm.userId != NEW.userId;  
END;


CREATE TRIGGER send_new_event_notification
AFTER INSERT ON eventsStatus
FOR EACH ROW
WHEN NEW.pending = 'pending'
BEGIN
    INSERT INTO notifications (userId, content, type, idRef)
    VALUES (
        NEW.userId,
        'New event',
        "e_ref",
        NEW.id
    );
END;


CREATE TRIGGER delete_notification_on_eventsStatus_update
AFTER UPDATE OF pending ON eventsStatus
FOR EACH ROW
WHEN NEW.pending IN ('completed','rejected')
BEGIN
    DELETE FROM notifications WHERE idRef = NEW.id;
END;


CREATE TRIGGER delete_completed_notification
AFTER INSERT ON notifications
FOR EACH ROW
BEGIN
    DELETE FROM notifications
    WHERE id = NEW.id
    AND NEW.type = 'g_ref'
    AND EXISTS (
        SELECT 1
        FROM groupMembers
        WHERE id = NEW.idRef
        AND pending = 'completed'
    );
END;
