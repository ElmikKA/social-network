package db

import (
	"fmt"
	"social-network/pkg/models"
)

func (s *Store) AddEvent(event models.Event) error {
	query := `INSERT INTO events (groupId,userId, title, description, time) VALUES (?,?,?,?,?)`
	_, err := s.Db.Exec(query, event.GroupId, event.UserId, event.Title, event.Description, event.Time)
	if err != nil {
		fmt.Println("err adding event", err)
		return err
	}
	return nil
}

func (s *Store) RespondEvent(userId, eventId int, response string) error {
	query := `UPDATE eventsStatus SET pending = ? WHERE userId = ? AND eventId = ?`
	_, err := s.Db.Exec(query, userId, response, eventId)
	if err != nil {
		fmt.Println("error responding to event")
		return err
	}
	return nil
}
