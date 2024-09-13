package db

import (
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

func (s *Store) AddGroupMember(group models.Group) error {
	query := `
	INSERT INTO groupMembers (userId, groupId) VALUES (?,?)
	`
	_, err := s.Db.Exec(query, group.UserId, group.Id)
	if err != nil {
		fmt.Println("err adding group member", err)
		return err
	}
	return nil
}
