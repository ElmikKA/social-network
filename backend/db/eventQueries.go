package db

import (
	"fmt"
	"social-network/pkg/models"
)

func (s *Store) AddEvent(event models.Event) error {
	query := `INSERT INTO events (groupId,userId, title, description, time) VALUES (?,?,?,?,?)`
	_, err := s.Db.Exec(query, event.GroupId, event.GroupId, event.Title, event.Description, event.Time)
	if err != nil {
		fmt.Println("err adding event", err)
		return err
	}
	return nil
}
