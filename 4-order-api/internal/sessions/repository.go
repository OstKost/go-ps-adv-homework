package sessions

import (
	"go-ps-adv-homework/pkg/db"
)

type SessionRepository struct {
	database *db.Db
}

func NewSessionRepository(db *db.Db) *SessionRepository {
	return &SessionRepository{
		database: db,
	}
}

func (repository *SessionRepository) GetSession(sessionStr string) (*Session, error) {
	var session Session
	result := repository.database.DB.
		Where("session = ?", sessionStr).
		First(&session)
	if result.Error != nil {
		return nil, result.Error
	}
	return &session, nil
}

func (repository *SessionRepository) CreateOrUpdateSession(payload *Session) (*Session, error) {
	var session Session
	result := repository.database.DB.
		Where(Session{Phone: payload.Phone}).
		Assign(Session{Session: payload.Session, Code: payload.Code}).
		FirstOrCreate(&session)
	if result.Error != nil {
		return nil, result.Error
	}
	return &session, nil
}

func (repository *SessionRepository) DeleteSession(id uint) error {
	result := repository.database.DB.Delete(&Session{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
