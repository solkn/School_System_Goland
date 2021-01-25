package repository

import (
	"database/sql"
	"errors"
	"github.com/Rob-a21/Cassiopeia/entity"
)

type PsqlNotificationRepositoryImpl struct {
	conn *sql.DB
}

const (
	layoutISO = "2006-01-02"
	layoutUS  = "January 2, 2006"
)

func NewPsqlNotificationRepositoryImpl(Conn *sql.DB) *PsqlNotificationRepositoryImpl {
	return &PsqlNotificationRepositoryImpl{conn: Conn}
}

func (ntf *PsqlNotificationRepositoryImpl) AddNotification(notf entity.Notification) error {

	_, err := ntf.conn.Exec("insert into notification (notifyname,message,notifieddate) values($1, $2,$3)", notf.NotifyName, notf.Message, notf.NotificationDate.Format(layoutUS))
	if err != nil {
		return errors.New("Insertion has failed")
	}

	return nil
}

func (ntf *PsqlNotificationRepositoryImpl) GetNotification() ([]entity.Notification, error) {

	rows, err := ntf.conn.Query("SELECT * FROM notification;")
	if err != nil {
		return nil, errors.New("Could not query the database")
	}
	defer rows.Close()

	notifications := []entity.Notification{}


	for rows.Next() {
		notification := entity.Notification{}
		err = rows.Scan(&notification.NotifyName, &notification.Message,&notification.NotificationDate)
		if err != nil {
			return nil, err
		}
		notifications = append(notifications, notification)
	}

	return notifications, err
}


