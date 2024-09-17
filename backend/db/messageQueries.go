package db

import (
	"fmt"
	"social-network/pkg/models"
)

func (s *Store) AddMessage(msg models.Message) error {
	query := `INSERT INTO messages (userId, message, receiverId, groupId) VALUES (?,?,?,?)`
	_, err := s.Db.Exec(query, msg.UserId, msg.Message, msg.ReceiverId, msg.GroupId)
	if err != nil {
		fmt.Println("error adding message", err)
		return err
	}
	return nil
}

func (s *Store) GetMessages(partner, groupId, userId int) ([]models.Message, error) {
	query := `
	SELECT 
	    m.userId, 
	    u.name AS name, 
	    m.message, 
	    m.receiverId, 
	    m.groupId 
	FROM 
	    messages m
	JOIN 
	    users u ON m.userId = u.id
	WHERE 
	    (m.userId = ? AND m.receiverId = ?) OR (m.userId = ? AND m.receiverId = ?) 
	AND 
	    m.groupId = ?
	ORDER BY 
	    m.createdAt;
	`

	rows, err := s.Db.Query(query, userId, partner, partner, userId, groupId)
	if err != nil {
		fmt.Println("error getting messages", err)
		return nil, err
	}
	defer rows.Close()

	var messages []models.Message
	for rows.Next() {
		var msg models.Message
		err := rows.Scan(&msg.UserId, &msg.Name, &msg.Message, &msg.ReceiverId, &msg.GroupId)
		if err != nil {
			fmt.Println("error scanning message", err)
			return nil, err
		}
		messages = append(messages, msg)
	}
	return messages, nil
}
