package db

import (
	"fmt"
)

func (s *Store) AddFollower(userId, follow int, pending string) (int, error) {
	// existsQuery := `SELECT COUNT(1) FROM followers WHERE userId = ? AND following = ?`
	// var count int
	// err := s.Db.QueryRow(existsQuery, userId, follow).Scan(&count)
	// if err != nil && err != sql.ErrNoRows {
	// 	// Handle error during the existence check
	// 	fmt.Println("error checking if follower exists", err)
	// 	return 0, err
	// }

	// // If the row exists, don't insert it
	// if count > 0 {
	// 	fmt.Println("follow already exists")
	// 	return 0, nil
	// }

	query := `INSERT INTO followers (userId, following, pending) VALUES (?,?,?)`
	result, err := s.Db.Exec(query, userId, follow, pending)
	if err != nil {
		fmt.Println("error adding following", err)
		return 0, err
	}
	newId, err := result.LastInsertId()
	if err != nil {
		fmt.Println("err getting last insert id follower", err)
		return 0, err
	}
	return int(newId), nil
}

func (s *Store) RespondFollow(userId, responseId int, answer string) error {
	query := `UPDATE followers SET pending = ? WHERE userId = ? AND following = ?`
	_, err := s.Db.Exec(query, answer, responseId, userId)
	if err != nil {
		fmt.Println("error responding followers", err)
		return err
	}
	return nil
}
