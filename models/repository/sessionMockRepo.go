package repository

import (
	"database/sql"
	"errors"
	"github.com/Rob-a21/Cassiopeia/entity"
	"github.com/Rob-a21/Cassiopeia/models"
)

type sessionMockRepo struct {
	conn *sql.DB
}

func NewsessionMockRepo(db *sql.DB)  models.SessionRepository {
	return &sessionMockRepo{conn: db}
}

func (sessionMockRepo) Session(sessionID string) (*entity.Session, error) {
	s := entity.SessionMock

	x := string(s.ID)

	if sessionID == x {
		return &entity.SessionMock, nil
	}

	return nil, nil
}

func (sessionMockRepo) StoreSession(session *entity.Session) (*entity.Session, error) {
	entity.SessionMock.ID = session.ID
	entity.SessionMock.Expires = session.Expires
	entity.SessionMock.SigningKey = session.SigningKey
	entity.SessionMock.UUID = session.UUID

	if session.ID != entity.SessionMock.ID{
		return nil,errors.New("Failed Registration")
	}
	return &entity.SessionMock,nil
}

func (sessionMockRepo) DeleteSession(sessionID string) (*entity.Session, error) {
	entity.SessionMock.ID = 0
	entity.SessionMock.Expires = 0
	entity.SessionMock.SigningKey = nil
	entity.SessionMock.UUID = ""

	x := string(entity.SessionMock.ID)

	if sessionID != x {
		return nil,errors.New("Failed Registration")
	}
	return &entity.SessionMock,nil

}


