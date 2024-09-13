package db

import (
	"fmt"
	"social-network/pkg/models"
)

func (s *Store) AddNotification(notification models.Notification) error {
	query := `
		INSERT INTO notifications (userId, content, type, idRef) VALUES (?,?,?,?)
	`
	_, err := s.Db.Exec(query, notification.UserId, notification.Content, notification.Type, notification.IdRef)
	if err != nil {
		fmt.Println("error adding notification")
		return err
	}
	return nil
}
