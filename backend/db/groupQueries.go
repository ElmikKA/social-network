package db

import (
	"database/sql"
	"fmt"
	"social-network/pkg/models"
)

func (s *Store) AddGroup(group models.Group) (int, error) {
	query := `
		INSERT INTO groups (userId, title, description) VALUES (?,?,?)
	`
	result, err := s.Db.Exec(query, group.UserId, group.Title, group.Description)
	if err != nil {
		fmt.Println("err adding group", err)
		return 0, err
	}
	newId, err := result.LastInsertId()
	if err != nil {
		fmt.Println("err getting last id addgroup", err)
		return 0, err
	}
	return int(newId), nil
}

func (s *Store) AddGroupMember(group models.Group) (int, error) {

	existsQuery := `SELECT COUNT(1) FROM groupMembers WHERE userId = ? AND groupId = ?`
	var count int
	err := s.Db.QueryRow(existsQuery, group.UserId, group.Id).Scan(&count)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println("error checking if follower exists", err)
		return 0, err
	}

	if count > 0 {
		fmt.Println("follow already exists")
		return 0, nil
	}
	query := `
	INSERT INTO groupMembers (userId, groupId) VALUES (?,?)
	`
	result, err := s.Db.Exec(query, group.UserId, group.Id)
	if err != nil {
		fmt.Println("err adding group member", err)
		return 0, err
	}
	newId, err := result.LastInsertId()
	if err != nil {
		fmt.Println("err getting lastid adding group member", err)
		return 0, err
	}

	return int(newId), nil
}

func (s *Store) GetOnlineGroupMembers(groupId int) ([]int, error) {

	fmt.Println(groupId)
	query := `
	SELECT u.id 
FROM groupMembers gm
JOIN users u ON gm.userId = u.id
WHERE gm.groupId = ?
AND u.online = 1;
`
	var response []int

	rows, err := s.Db.Query(query, groupId)
	if err != nil {
		fmt.Println("error getting online group members", err)
		return response, err
	}

	defer rows.Close()

	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			fmt.Println("err rows.next groupmembers", err)
			return response, err
		}
		response = append(response, id)
	}
	return response, nil
}

func (s *Store) GetGroupMembers(groupId int) ([]models.GroupMember, error) {
	query := `
	SELECT 
    gm.userId,
    u.name,
    u.avatar,
    gm.role
FROM 
    groupMembers gm
JOIN 
    users u
ON 
    gm.userId = u.id
WHERE 
    gm.groupId = ?  
AND 
    gm.pending = 'completed';

`
	var members []models.GroupMember
	rows, err := s.Db.Query(query, groupId)
	if err != nil {
		fmt.Println("error getting group member info", err)
		return members, err
	}

	defer rows.Close()
	for rows.Next() {
		var member models.GroupMember
		err := rows.Scan(&member.Id, &member.Name, &member.Avatar, &member.Role)
		if err != nil {
			fmt.Println("error scanning member info", err)
			return members, err
		}
		members = append(members, member)
	}
	return members, nil
}

func (s *Store) GetGroupEvents(groupId, userId int) ([]models.GroupEvents, error) {
	query := `
	SELECT 
	    e.id AS Id,
	    e.userId AS OwnerId,
	    e.groupId AS GroupId,
	    e.title AS Title,
	    e.description AS Description,
	    e.time AS Time,
	    es.pending AS Status,
	    es.role AS Role,
	    es.id AS IdRef
	FROM 
	    events e
	LEFT JOIN 
	    eventsStatus es 
	ON 
	    e.id = es.eventId
	AND 
	    es.userId = ?
	WHERE 
	    e.groupId = ?;
	`

	var events []models.GroupEvents

	rows, err := s.Db.Query(query, userId, groupId)
	if err != nil {
		if err == sql.ErrNoRows {
			return events, nil
		}
		fmt.Println("error getting group event info", err)
		return events, err
	}

	defer rows.Close()

	for rows.Next() {
		var event models.GroupEvents
		err := rows.Scan(&event.Id, &event.OwnerId, &event.GroupId, &event.Title, &event.Description, &event.Time, &event.Status, &event.Role, &event.IdRef)
		if err != nil {
			fmt.Println("error scanning event info", err)
			return events, err
		}
		events = append(events, event)
	}
	return events, nil
}

func (s *Store) GetIsPartOfGroup(groupId, userId int) (bool, error) {
	query := `SELECT COUNT(*) FROM groupMembers WHERE userId = ? AND groupId = ? AND pending = 'completed'`
	var count int
	err := s.Db.QueryRow(query, userId, groupId).Scan(&count)
	if err != nil {
		fmt.Println("error getting part of group", err)
		return false, err
	}
	return count > 0, nil
}

func (s *Store) GetAllGroups() ([]models.Group, error) {
	query := `SELECT * FROM groups`
	rows, err := s.Db.Query(query)
	var groups []models.Group
	if err != nil {
		fmt.Println("err getting all groups", err)
		return groups, err
	}
	defer rows.Close()
	var group models.Group
	for rows.Next() {
		err := rows.Scan(&group.Id, &group.UserId, &group.Title, &group.Description)
		if err != nil {
			fmt.Println("err scanning all groups", err)
			return groups, err
		}
		groups = append(groups, group)
	}
	return groups, nil
}

func (s *Store) GetGroup(groupId int) (models.Group, error) {
	query := `SELECT * FROM groups WHERE id = ?`
	var group models.Group
	err := s.Db.QueryRow(query, groupId).Scan(&group.Id, &group.UserId, &group.Title, &group.Description)
	if err != nil {
		fmt.Println("err getting group", err)
		return group, err
	}
	return group, nil
}

func (s *Store) GetGroupJoinStatus(groupId, userId int) (string, error) {
	query := `SELECT pending FROM groupMembers WHERE groupId = ? AND userId = ?`
	var status string
	err := s.Db.QueryRow(query, groupId, userId).Scan(&status)
	if err == sql.ErrNoRows {
		return "", nil
	}
	if err != nil {
		fmt.Println("err getting group join status", err)
		return status, err
	}
	return status, nil

}
