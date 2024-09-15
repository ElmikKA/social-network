package db

import (
	"fmt"
	"social-network/pkg/models"
)

func (s *Store) AddPost(post models.Post) error {
	query := `
		INSERT INTO posts (userId, groupId, creator, title, content, avatar, privacy) VALUES (?,?,?,?,?,?,?)
	`
	_, err := s.Db.Exec(query, post.UserId, post.GroupId, post.Creator, post.Title, post.Content, post.Avatar, post.Privacy)
	if err != nil {
		fmt.Println("error adding new post", err)
		return err
	}
	return nil
}

func (s *Store) GetPost(id int) (models.Post, error) {
	query := `
		SELECT * FROM posts WHERE id =?
	`
	var post models.Post
	err := s.Db.QueryRow(query, id).Scan(&post.Id, &post.UserId, &post.GroupId, &post.Creator, &post.Title, &post.Content, &post.Avatar, &post.CreatedAt, &post.Privacy)
	if err != nil {
		fmt.Println("error getting one post", err)
		return models.Post{}, err
	}
	return post, nil
}
func (s *Store) GetAllUserPosts(userId int) ([]models.Post, error) {

	query := `
		SELECT * FROM posts WHERE id =?
	`
	rows, err := s.Db.Query(query, userId)
	if err != nil {
		fmt.Println("error getting all posts", err)
		return []models.Post{}, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		rows.Scan(&post.Id, &post.UserId, &post.GroupId, &post.Creator, &post.Title, &post.Content, &post.Avatar, &post.CreatedAt, &post.Privacy)
		posts = append(posts, post)
	}
	return posts, nil
}

func (s *Store) GetAllNormalPosts() ([]models.Post, error) {
	query := `
		SELECT * FROM POSTS WHERE groupId = 0
	`
	rows, err := s.Db.Query(query)
	if err != nil {
		fmt.Println("error getting all posts", err)
		return []models.Post{}, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		rows.Scan(&post.Id, &post.UserId, &post.GroupId, &post.Creator, &post.Title, &post.Content, &post.Avatar, &post.CreatedAt, &post.Privacy)
		posts = append(posts, post)
	}
	return posts, nil
}

func (s *Store) GetAllGroupPosts(groupId int) ([]models.Post, error) {
	query := `
		SELECT * FROM POSTS WHERE groupId = ?
	`
	rows, err := s.Db.Query(query, groupId)
	if err != nil {
		fmt.Println("error getting all posts", err)
		return []models.Post{}, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		rows.Scan(&post.Id, &post.UserId, &post.GroupId, &post.Creator, &post.Title, &post.Content, &post.Avatar, &post.CreatedAt, &post.Privacy)
		posts = append(posts, post)
	}
	return posts, nil
}
