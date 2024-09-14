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

func (s *Store) GetOnlineGroupMembers(userId int) ([]int, error) {
	query := `SELECT u.id 
FROM groupMembers gm
JOIN users u ON gm.userId = u.id
WHERE gm.groupId = (SELECT groupId FROM groupMembers WHERE userId = ?)
AND u.id != ? 
AND u.online = 1;
`
	var response []int

	rows, err := s.Db.Query(query, userId, userId)
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