package models

import (
	"net/http"
)

// add every function that uses db here
type UserStore interface {
	CheckUserExists(user Users) (bool, error)
	AddUser(user Users) error
	CheckLogin(credentials LoginCredentials) (bool, int, error)
	GetUserFromCookie(r *http.Request) (Users, error)
	GetUser(userId int) (Users, error)
	GetAllUsers() ([]Users, error)
	GetSessionByCookie(cookie string) (Session, error)
	CreateSession(session Session) error
	DeleteSession(id int) error
	ExtendSessionDate(cookie string) error
	ResetOnline()
	GoOffline(id int) error
	GoOnline(id int) error
	AddPost(post Post) error
	GetPost(id int) (Post, error)
	GetAllUserPosts(userId int) ([]Post, error)
	GetAllNormalPosts() ([]Post, error)
	GetAllGroupPosts(groupId int) ([]Post, error)
	GetGroupMembers(groupId int) ([]GroupMember, error)
	GetIsPartOfGroup(groupId, userId int) (bool, error)
	GetGroupEvents(groupId, userId int) ([]GroupEvents, error)
	AddComment(comment Comment) error
	GetComments(postId int) ([]Comment, error)
	AddFollower(userId, following int, pending string) (int, error)
	CheckUserPrivacyStatus(userId int) (string, error)
	AddNotification(notification Notification) error
	RespondNotification(response NotificationResponse) error
	AddGroup(group Group) (int, error)
	AddGroupMember(group Group) (int, error)
	AddEvent(event Event) error
	RespondEvent(userId, eventId int, answer string) error
	AddMessage(msg Message) error
	GetOnlineGroupMembers(userId int) ([]int, error)
	GetNotifications(userId int) ([]Notification, error)
	GetContacts(userId int) ([]Contacts, error)
	GetGroupChats(userId int) ([]GroupChats, error)
	GetMessages(userId, groupId int) ([]Message, error)
	IsFollowing(userId, followee int) (string, error)
	GetAllGroups() ([]Group, error)
	GetGroup(groupId int) (Group, error)
}
