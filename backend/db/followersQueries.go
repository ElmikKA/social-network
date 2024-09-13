package db

import "fmt"

func (s *Store) AddFollower(userId, follow int, pending string) (int, error) {
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
