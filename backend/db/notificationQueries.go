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

func (s *Store) RespondNotification(response models.NotificationResponse) error {
	query := ``
	switch response.Type {
	case "g_ref":
		query = `UPDATE groupMembers SET pending = ? WHERE id = ?`
		fmt.Println("g_ref")
	case "f_ref":
		query = `UPDATE followers SET pending = ? WHERE id = ?`
	case "e_ref":
		query = `UPDATE eventsStatus SET pending = ? WHERE id = ?`
	}
	_, err := s.Db.Exec(query, response.Response, response.IdRef)
	if err != nil {
		fmt.Println("err responding notification")
		return err
	}
	return nil
}

func (s *Store) GetNotifications(userId int) ([]models.Notification, error) {
	query := `SELECT * FROM notifications WHERE userId = ?`
	rows, err := s.Db.Query(query, userId)
	var notifications []models.Notification
	if err != nil {
		fmt.Println("error getting notifications", err)
		return notifications, err
	}
	defer rows.Close()
	for rows.Next() {
		var noti models.Notification
		err := rows.Scan(&noti.Id, &noti.UserId, &noti.Content, &noti.Type, &noti.IdRef, &noti.CreatedAt)
		if err != nil {
			fmt.Println("err scanning notifications", err)
			return notifications, err
		}
		notifications = append(notifications, noti)
	}
	return notifications, nil
}
