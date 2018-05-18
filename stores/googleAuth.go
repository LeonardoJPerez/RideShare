package stores

import (
	"time"

	"github.com/RideShare-Server/models"

	"github.com/jinzhu/gorm"
	"github.com/juju/errors"
	"github.com/volatiletech/authboss"
)

// GoogleAuthStore is responsible for persisting and getting GoogleAuthData objects
type GoogleAuthStore struct {
	BaseStore
}

// AuthResult is a struct that mirrors the GoogleAuthData models, without the ID field
// which causes issues with AuthBoss' unbind function
type AuthResult struct {
	// AuthBoss Key
	Key string `json:"key,omitempty"`

	// Auth
	Email string `json:"email,omitempty"`

	// OAuth2
	Oauth2Uid      string    `json:"oauth2_uid,omitempty"`
	Oauth2Provider string    `json:"oauth2_provider,omitempty"`
	Oauth2Token    string    `json:"oauth2_token,omitempty"`
	Oauth2Expiry   time.Time `json:"oauth2_expiry,omitempty"`
}

// NewGoogleAuthStore create a new GoogleAuthStore with the database
func NewGoogleAuthStore(db *gorm.DB) *GoogleAuthStore {
	GoogleAuthStore := new(GoogleAuthStore)
	GoogleAuthStore.database = db
	return GoogleAuthStore
}

// Put updates the GoogleAuthData record in the database
func (s *GoogleAuthStore) Put(key string, attr authboss.Attributes) error {
	var authDataAttributes models.GoogleAuth
	if err := attr.Bind(&authDataAttributes, true); err != nil {
		return errors.Trace(err)
	}

	// Try and find the GoogleAuthData entry and update it, or create a new one
	authDataAttributes.Key = key
	err := s.database.
		Where(models.GoogleAuth{Key: key}).
		Assign(authDataAttributes).
		FirstOrCreate(&models.GoogleAuth{}).Error
	if err != nil {
		return errors.Trace(err)
	}
	return nil
}

// Get fetches the GoogleAuthData record from the database
func (s *GoogleAuthStore) Get(key string) (result interface{}, err error) {
	authData := &models.GoogleAuth{}
	err = s.database.Where("key = ?", key).First(authData).Error
	if err != nil {
		return nil, errors.Trace(err)
	}

	interfaceResult := &AuthResult{
		Email:          authData.Email,
		Key:            authData.Key,
		Oauth2Expiry:   authData.Oauth2Expiry,
		Oauth2Provider: authData.Oauth2Provider,
		Oauth2Token:    authData.Oauth2Token,
		Oauth2Uid:      authData.Oauth2Uid,
	}

	return interfaceResult, nil
}

// PutOAuth updates the GoogleAuthData record in the database
func (s *GoogleAuthStore) PutOAuth(uid string, provider string, attr authboss.Attributes) error {
	var authDataAttributes models.GoogleAuth
	if err := attr.Bind(&authDataAttributes, true); err != nil {
		return errors.Trace(err)
	}

	// Try and find the GoogleAuthData entry and update it, or create a new one
	authDataAttributes.Key = uid + provider
	err := s.database.
		Where(models.GoogleAuth{Oauth2Uid: uid, Oauth2Provider: provider}).
		Assign(authDataAttributes).
		FirstOrCreate(&models.GoogleAuth{}).Error
	if err != nil {
		return errors.Trace(err)
	}
	return nil
}

// GetOAuth fetches the GoogleAuthData record from the database
func (s *GoogleAuthStore) GetOAuth(uid string, provider string) (result interface{}, err error) {
	authData := &models.GoogleAuth{}
	key := uid + provider
	err = s.database.Where("key = ?", key).First(authData).Error
	if err != nil {
		return nil, errors.Trace(err)
	}

	interfaceResult := &AuthResult{
		Email:          authData.Email,
		Key:            authData.Key,
		Oauth2Expiry:   authData.Oauth2Expiry,
		Oauth2Provider: authData.Oauth2Provider,
		Oauth2Token:    authData.Oauth2Token,
		Oauth2Uid:      authData.Oauth2Uid,
	}
	return interfaceResult, nil
}
