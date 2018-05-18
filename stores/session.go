package stores

import (
	"math/rand"
	"time"

	"github.com/RideShare-Server/db"
	"github.com/RideShare-Server/log"
	"github.com/RideShare-Server/models"
	"github.com/jinzhu/gorm"
	"github.com/juju/errors"
)

var (
	s1 = rand.NewSource(time.Now().UnixNano())
	r1 = rand.New(s1)
)

// SessionStore :
type SessionStore struct {
	BaseStore
}

// NewSessionStore :
func NewSessionStore(db *gorm.DB) *SessionStore {
	store := new(SessionStore)
	store.database = db

	return store
}

// Get attempts to retrieve an existing session record given a UserID and SessionID.
func (s *SessionStore) Get(userID uint, sessionID string) (*models.UserSession, error) {
	if userID == 0 {
		err := errors.New("user id cannot be empty")
		log.Error(log.StoreLayerTopic, err)
		return nil, errors.Trace(err)
	}

	if sessionID == "" {
		err := errors.New("session id cannot be empty")
		log.Error(log.StoreLayerTopic, err)
		return nil, errors.Trace(err)
	}

	session := &models.UserSession{}
	err := s.database.
		Where("user_id = ? AND session_id = ?", userID, sessionID).
		Order("signed_in desc").
		First(session).Error

	if err == nil {
		return session, nil
	}

	if db.CheckNotFoundErr(err) != nil {
		log.Error(log.StoreLayerTopic, err)
		return nil, errors.Trace(err)
	}

	return nil, nil
}

// Validate attempts to get an existing session record given a UserID and SessionID, then validates it exist.
func (s *SessionStore) Validate(userID uint, sessionID string) (bool, error) {
	session, err := s.Get(userID, sessionID)
	if err != nil {
		return false, err
	}

	return session != nil, nil
}

// Insert a new Session record in UserSessions table.
func (s *SessionStore) Insert(session *models.UserSession) (*models.UserSession, error) {
	err := s.database.Create(session).Error
	if err != nil {
		log.Error(log.StoreLayerTopic, err)
		return nil, errors.Trace(err)
	}

	return session, nil
}
