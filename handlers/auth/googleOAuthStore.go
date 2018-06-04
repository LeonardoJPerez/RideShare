package auth

import (
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/juju/errors"
	"github.com/volatiletech/authboss"
)

// GoogleOAuthStore is responsible for persisting and getting GoogleOAuthData objects
type GoogleOAuthStore struct {
	database *gorm.DB
}

// OAuth2Result is a struct that mirrors the GoogleOAuth model, without the ID field
// which causes issues with AuthBoss' unbind function
type OAuth2Result struct {
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

// NewGoogleOAuthStore create a new GoogleOAuthStore with the database
func NewGoogleOAuthStore(database *gorm.DB) *GoogleOAuthStore {
	store := new(GoogleOAuthStore)
	store.database = database

	return store
}

// Store Methods.

// Put updates the GoogleOAuthData record in the database
func (store *GoogleOAuthStore) Put(key string, attr authboss.Attributes) error {
	var authDataAttributes GoogleOAuth
	if err := attr.Bind(&authDataAttributes, true); err != nil {
		return errors.Trace(err)
	}

	// Try and find the GoogleOAuthData entry and update it, or create a new one
	var foundAuthData GoogleOAuth
	authDataAttributes.Key = key
	err := store.database.
		Where(GoogleOAuth{Key: key}).
		Assign(authDataAttributes).
		FirstOrCreate(&foundAuthData).Error
	if err != nil {
		return errors.Trace(err)
	}
	return nil
}

// Get fetches the GoogleOAuthData record from the database
func (store *GoogleOAuthStore) Get(key string) (result interface{}, err error) {
	var foundAuthData GoogleOAuth
	if err := store.database.
		Where("google_auth_data.key = ?", key).
		First(&foundAuthData).Error; err != nil {
		return nil, errors.Trace(err)
	}

	interfaceResult := &OAuth2Result{
		Email:          foundAuthData.Email,
		Key:            foundAuthData.Key,
		Oauth2Expiry:   foundAuthData.Oauth2Expiry,
		Oauth2Provider: foundAuthData.Oauth2Provider,
		Oauth2Token:    foundAuthData.Oauth2Token,
		Oauth2Uid:      foundAuthData.Oauth2Uid,
	}
	return interfaceResult, nil
}

// PutOAuth updates the GoogleOAuth record in the database
func (store *GoogleOAuthStore) PutOAuth(uid string, provider string, attr authboss.Attributes) error {
	var authDataAttributes GoogleOAuth
	if err := attr.Bind(&authDataAttributes, true); err != nil {
		return errors.Trace(err)
	}

	// Check for GawkBox and Mesh emails
	isGawkboxEmail := strings.Contains(authDataAttributes.Email, "@gawkbox.com")
	isMeshEmail := strings.Contains(authDataAttributes.Email, "@meshstudio.io")
	if !isGawkboxEmail && !isMeshEmail {
		return errors.New("Invalid domain name")
	}

	// Try and find the GoogleOAuthData entry and update it, or create a new one
	var foundAuthData GoogleOAuth
	authDataAttributes.Key = uid + provider
	err := store.database.
		Where(GoogleOAuth{Oauth2Uid: uid, Oauth2Provider: provider}).
		Assign(authDataAttributes).
		FirstOrCreate(&foundAuthData).Error
	if err != nil {
		return errors.Trace(err)
	}
	return nil
}

// GetOAuth fetches the GoogleOAuth record from the database
func (store *GoogleOAuthStore) GetOAuth(uid string, provider string) (result interface{}, err error) {
	var foundAuthData GoogleOAuth
	key := uid + provider
	if err := store.database.
		Where("google_auth_data.key = ?", key).
		First(&foundAuthData).Error; err != nil {
		return nil, errors.Trace(err)
	}

	interfaceResult := &OAuth2Result{
		Email:          foundAuthData.Email,
		Key:            foundAuthData.Key,
		Oauth2Expiry:   foundAuthData.Oauth2Expiry,
		Oauth2Provider: foundAuthData.Oauth2Provider,
		Oauth2Token:    foundAuthData.Oauth2Token,
		Oauth2Uid:      foundAuthData.Oauth2Uid,
	}
	return interfaceResult, nil
}
