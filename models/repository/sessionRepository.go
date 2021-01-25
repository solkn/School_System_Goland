package repository

import (
	"database/sql"

	"github.com/Rob-a21/Cassiopeia/entity"
)

type SessionRepo struct {
	conn *sql.DB
}

func NewSessionRepo(db *sql.DB) *SessionRepo {
	return &SessionRepo{conn: db}
}

func (sr *SessionRepo) Session(sessionID string) (*entity.Session, error) {

	row := sr.conn.QueryRow("SELECT * FROM session WHERE id = $1", sessionID)

	session := entity.Session{}

	err := row.Scan(&session.ID, &session.Expires, &session.UUID, &session.SigningKey)
	if err != nil {
		panic(err)
	}

	return &session, err
}

// StoreSession stores a given session
func (sr *SessionRepo) StoreSession(session *entity.Session) (*entity.Session, error) {

	_, err := sr.conn.Exec("insert into session (id,uuid,expires,signingkey) values($1, $2,$3,&4)", session.ID, session.UUID, session.Expires, session.SigningKey)
	if err != nil {

		panic(err)

	}

	return nil, err
}

// DeleteSession deletes a given session
func (sr *SessionRepo) DeleteSession(sessionID string) (*entity.Session, error) {

	_, err := sr.conn.Exec("DELETE FROM session WHERE id=$1", sessionID)

	if err != nil {
		return nil, err
	}

	return nil, err
}
