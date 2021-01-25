package service

import (
	"github.com/Rob-a21/Cassiopeia/entity"
	"github.com/Rob-a21/Cassiopeia/models"
)

type NotificationServiceImpl struct {
	notificationRepository models.NotificationRepository
}

func NewNotificationServiceImpl(nrpo models.NotificationRepository) *NotificationServiceImpl {
	return &NotificationServiceImpl{notificationRepository: nrpo}
}

func (ntf *NotificationServiceImpl) AddNotification(notf entity.Notification) error {

	err := ntf.notificationRepository.AddNotification(notf)

	if err != nil {
		return err
	}

	return nil
}

func (ss *NotificationServiceImpl) GetNotification() ([]entity.Notification, error) {

	notn, err := ss.notificationRepository.GetNotification()

	if err != nil {
		return nil, err
	}

	return notn, nil
}

