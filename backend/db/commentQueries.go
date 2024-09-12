package db

import (
	"fmt"
	"social-network/pkg/models"
)

func (s *Store) AddComment(comment models.Comment) error {
	query := `
		INSERT INTO comments (userId, postId, creator, content, avatar) 
		VALUES (?,?,?,?,?)
	`
	_, err := s.Db.Exec(query, comment.UserId, comment.PostId, comment.Creator, comment.Content, comment.Avatar)
	if err != nil {
		fmt.Println("error adding comment", err)
		return err
	}
	return nil
}

func (s *Store) GetComments(postId int) ([]models.Comment, error) {
	query := `
		SELECT * FROM comments WHERE postId = ?
	`
	rows, err := s.Db.Query(query, postId)
	if err != nil {
		fmt.Println("error getting comments", err)
		return []models.Comment{}, err
	}
	defer rows.Close()
	var comments []models.Comment
	for rows.Next() {
		var comment models.Comment
		err := rows.Scan(&comment.Id, &comment.UserId, &comment.PostId, &comment.Creator, &comment.Content, &comment.CreatedAt, &comment.Avatar)
		if err != nil {
			fmt.Println("err scanning comments rows", err)
			return []models.Comment{}, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}
