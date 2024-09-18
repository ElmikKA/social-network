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




-- CREATE TRIGGER set_groupMembers_role_and_pending_and_create_notification
-- AFTER INSERT ON groupMembers
-- FOR EACH ROW
-- BEGIN
--     UPDATE groupMembers
--     SET role = CASE
--         WHEN (SELECT userId FROM groups WHERE id = NEW.groupId) = NEW.userId
--         THEN 'owner'
--         ELSE 'member'
--     END,
--     pending = CASE
--         WHEN (SELECT userId FROM groups WHERE id = NEW.groupId) = NEW.userId
--         THEN 'completed'
--         ELSE 'pending'
--     END
--     WHERE id = NEW.id;

--     INSERT INTO notifications (userId, content, type, idRef)
--     SELECT 
--         (SELECT userId FROM groups WHERE id = NEW.groupId), 
--         (SELECT name FROM users WHERE id = NEW.userId) || ' has sent a request to join your group ' || 
--         (SELECT title FROM groups WHERE id = NEW.groupId),  
--         'g_ref', 
--         NEW.id  
--     FROM groupMembers
--     WHERE id = NEW.id AND pending = 'pending';  
-- END;



CREATE TRIGGER set_groupMembers_role_and_pending_and_create_notification
AFTER INSERT ON groupMembers
FOR EACH ROW
BEGIN
    -- Update the role and pending status based on whether the user is the owner
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

    -- Send a notification based on the value of 'pending' and 'invitee'
    INSERT INTO notifications (userId, content, type, idRef)
    SELECT
        CASE
            WHEN NEW.invitee = 0 AND NEW.pending = 'pending' THEN (SELECT userId FROM groups WHERE id = NEW.groupId)
            WHEN NEW.invitee = 1 THEN NEW.userId
            ELSE NULL
        END AS userId,
        CASE
            WHEN NEW.invitee = 0 AND NEW.pending = 'pending' THEN (SELECT name FROM users WHERE id = NEW.userId) || ' has sent a request to join your group ' || (SELECT title FROM groups WHERE id = NEW.groupId)
            WHEN NEW.invitee = 1 THEN 'You have been added to group ' || (SELECT title FROM groups WHERE id = NEW.groupId)
            ELSE NULL
        END AS content,
        CASE
            WHEN NEW.invitee = 0 AND NEW.pending = 'pending' THEN 'g_ref'
            WHEN NEW.invitee = 1 THEN 'gi_ref'
            ELSE NULL
        END AS type,
        NEW.id AS idRef
    WHERE
        (NEW.invitee = 0 AND NEW.pending = 'pending')
        OR
        NEW.invitee = 1;
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
