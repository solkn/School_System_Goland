package service

import (
	"github.com/Rob-a21/Cassiopeia/entity"
	"github.com/Rob-a21/Cassiopeia/models"
)

// SessionServiceImpl implements user.SessionService interface
type SessionServiceImpl struct {
	sessionRepo models.SessionRepository
}
//
//func b(sessRepository models.SessionRepository) models.SessionService {
//	return &SessionServiceImpl{sessionRepo: sessRepository}
//}
func New(sessRepository models.SessionRepository) models.SessionService {
	return &SessionServiceImpl{sessionRepo: sessRepository}
}
// Session returns a given stored session
func (ss *SessionServiceImpl) Session(sessionID string) (*entity.Session, error) {
	sess, err := ss.sessionRepo.Session(sessionID)
	if err != nil {
		return nil, err
	}
	return sess, err
}

// StoreSession stores a given session
func (ss *SessionServiceImpl) StoreSession(session *entity.Session) (*entity.Session, error) {
	sess, err := ss.sessionRepo.StoreSession(session)
	if err != nil {
		return nil, err
	}
	return sess, err
}

// DeleteSession deletes a given session
func (ss *SessionServiceImpl) DeleteSession(sessionID string) (*entity.Session, error) {
	sess, err := ss.sessionRepo.DeleteSession(sessionID)
	if err != nil {
		return nil, err
	}
	return sess, err
}