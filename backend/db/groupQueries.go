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
