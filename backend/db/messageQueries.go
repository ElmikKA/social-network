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
