package db

import (
	"database/sql"
	"fmt"
	"social-network/pkg/models"
)

func (s *Store) AddFollower(userId, follow int, pending string) (int, error) {
	existsQuery := `SELECT COUNT(1) FROM followers WHERE userId = ? AND following = ?`
	var count int
	err := s.Db.QueryRow(existsQuery, userId, follow).Scan(&count)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println("error checking if follower exists", err)
		return 0, err
	}

	if count > 0 {
		fmt.Println("follow already exists")
		return 0, nil
	}

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

func (s *Store) GetContacts(userId int) ([]models.Contacts, error) {
	query := `
	SELECT 
	    u.id, 
	    u.name, 
	    u.avatar,
	    CASE 
	        WHEN f.userId = ? THEN 'following'
	        WHEN f.following = ? THEN 'followee'
	    END AS type
	FROM followers f
	JOIN users u ON (f.following = u.id OR f.userId = u.id)
	WHERE (f.userId = ? OR f.following = ?) 
	AND f.pending = 'completed'
	AND u.id != ?;
	`

	rows, err := s.Db.Query(query, userId, userId, userId, userId, userId)
	if err != nil {
		fmt.Println("error getting contacts", err)
		return nil, err
	}
	defer rows.Close()

	var contacts []models.Contacts
	for rows.Next() {
		var contact models.Contacts
		err := rows.Scan(&contact.Id, &contact.Name, &contact.Avatar, &contact.Type)
		if err != nil {
			fmt.Println("error scanning contact info", err)
			return nil, err
		}
		contacts = append(contacts, contact)
	}

	return contacts, nil
}

func (s *Store) IsFollowing(userId, followee int) (string, error) {
	query := `SELECT pending FROM followers WHERE (userId = ? AND following = ?) OR (userId = ? AND following = ?)`
	var pending string
	err := s.Db.QueryRow(query, userId, followee, followee, userId).Scan(&pending)
	if err != nil {
		fmt.Println("err getting isfollowing", err)
		return pending, err
	}
	return pending, nil
}
