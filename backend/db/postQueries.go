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

func (s *Store) GetAllUserPosts(userId, creator int) ([]models.Post, error) {
	query := `
		SELECT p.id, p.userId, p.groupId, p.creator, p.title, p.content, p.avatar, p.createdAt, p.privacy
		FROM posts p
		WHERE p.userId = ?
	`
	rows, err := s.Db.Query(query, creator)
	if err != nil {
		fmt.Println("error getting all posts", err)
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.Id, &post.UserId, &post.GroupId, &post.Creator, &post.Title, &post.Content, &post.Avatar, &post.CreatedAt, &post.Privacy)
		if err != nil {
			fmt.Println("error scanning post info", err)
			return nil, err
		}

		if userId == creator {
			post.CanSee = true
		} else if post.Privacy == "private" {
			followingQuery := `
				SELECT COUNT(1)
				FROM followers
				WHERE userId = ? AND following = ? AND pending = 'completed'
			`
			var count int
			err := s.Db.QueryRow(followingQuery, userId, post.UserId).Scan(&count)
			if err != nil {
				fmt.Println("error checking following status", err)
				return nil, err
			}
			post.CanSee = count > 0
		} else {
			post.CanSee = true
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (s *Store) GetAllNormalPosts() ([]models.Post, error) {
	query := `
		SELECT * FROM POSTS WHERE groupId = 0 ORDER BY id DESC
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

func (s *Store) GetAllNormalPostsPrivacy(userId int) ([]models.Post, error) {
	query := `
        SELECT * FROM POSTS WHERE groupId = 0 ORDER BY id DESC
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
		err := rows.Scan(&post.Id, &post.UserId, &post.GroupId, &post.Creator, &post.Title, &post.Content, &post.Avatar, &post.CreatedAt, &post.Privacy)
		if err != nil {
			fmt.Println("error scanning row", err)
			return []models.Post{}, err
		}

		if post.UserId == userId {
			post.CanSee = true
		} else if post.Privacy == "public" {
			post.CanSee = true
		} else {
			followingQuery := `
                SELECT COUNT(*) FROM followers 
                WHERE userId = ? AND following = ? AND pending = 'completed'
            `
			var count int
			err = s.Db.QueryRow(followingQuery, userId, post.UserId).Scan(&count)
			if err != nil {
				fmt.Println("error checking if user follows post creator", err)
				return []models.Post{}, err
			}
			post.CanSee = (count > 0)
		}

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
