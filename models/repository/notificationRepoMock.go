package repository

import (
	"database/sql"

	"github.com/Rob-a21/Cassiopeia/entity"
	"github.com/Rob-a21/Cassiopeia/models"
)

type NotificationRepoMock struct {
	conn *sql.DB
}

func NewNotificationRepoMock(db *sql.DB) models.NotificationRepository {
	return &NotificationRepoMock{conn: db}
}

func (aRepo *NotificationRepoMock) GetNotification() ([]entity.Notification, error) {
	posts := []entity.Notification{entity.NotificationMock}

	return posts, nil
}

func (aRepo *NotificationRepoMock) AddNotification(notif entity.Notification) error {
	notif = entity.NotificationMock

	return nil
}

//func (aRepo *NotificationRepoMock) Notifications() ([]entity.Notification, []error) {
//
//	posts := []entity.Notification{entity.NotificationMock}
//
//	return posts, nil
//}
//
//func (aRepo *NotificationRepoMock) Notification(id uint) (*entity.Notification, []error) {
//
//	if id == 1 {
//		return &entity.NotificationMock, nil
//	}
//
//	return nil, nil
//
//}
//
//func (aRepo *NotificationRepoMock) UpdateNotification(ass *entity.Notification) (*entity.Notification, []error) {
//
//	pst := entity.NotificationMock
//
//	return &pst, nil
//}
//
//func (aRepo *NotificationRepoMock) DeleteNotification(id uint) (*entity.Notification, []error) {
//	post := entity.NotificationMock
//	if id != 1 {
//		return nil, []error{errors.New("post not found")}
//	}
//
//	return &post, nil
//}
