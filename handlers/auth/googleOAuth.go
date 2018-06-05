package auth

import (
	"os"
	"time"

	"github.com/RideShare-Server/utils"
	"github.com/jinzhu/gorm"
	"github.com/juju/errors"
	"github.com/volatiletech/authboss"
)

// GoogleAuthStore is responsible for persisting and getting GoogleAuthData objects
type (
	GoogleAuthStore struct {
		database *gorm.DB
	}

	// OAuth2Result is a struct that mirrors the GoogleAuth model, without the ID field
	// which causes issues with AuthBoss' unbind function
	OAuth2Result struct {
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

	// GoogleAuth represents the data we get from Google in the auth process
	GoogleAuth struct {
		ID        uint      `gorm:"primary_key" json:"id,omitempty"`
		CreatedAt time.Time `gorm:"" json:"created_at,omitempty"`
		UpdatedAt time.Time `gorm:"" json:"updated_at,omitempty"`

		// AuthBoss Key
		Key string `gorm:"" json:"key,omitempty"`

		// Auth
		Email string `gorm:"" json:"email,omitempty"`

		// OAuth2
		Oauth2Uid      string    `gorm:"" json:"oauth2_uid,omitempty"`
		Oauth2Provider string    `gorm:"" json:"oauth2_provider,omitempty"`
		Oauth2Token    string    `gorm:"" json:"oauth2_token,omitempty"`
		Oauth2Expiry   time.Time `gorm:"" json:"oauth2_expiry,omitempty"`
	}
)

// NewGoogleAuthStore create a new GoogleAuthStore with the database
func NewGoogleAuthStore(database *gorm.DB) *GoogleAuthStore {
	store := new(GoogleAuthStore)
	store.database = database

	store.migrate()

	return store
}

// Store Methods.

// Put updates the GoogleAuthData record in the database
func (store *GoogleAuthStore) Put(key string, attr authboss.Attributes) error {
	var authDataAttributes GoogleAuth
	if err := attr.Bind(&authDataAttributes, true); err != nil {
		return errors.Trace(err)
	}

	// Try and find the GoogleAuthData entry and update it, or create a new one
	var foundAuthData GoogleAuth
	authDataAttributes.Key = key
	err := store.database.
		Where(GoogleAuth{Key: key}).
		Assign(authDataAttributes).
		FirstOrCreate(&foundAuthData).Error
	if err != nil {
		return errors.Trace(err)
	}
	return nil
}

// Get fetches the GoogleAuthData record from the database
func (store *GoogleAuthStore) Get(key string) (result interface{}, err error) {
	var foundAuthData GoogleAuth
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

// PutOAuth updates the GoogleAuth record in the database
func (store *GoogleAuthStore) PutOAuth(uid string, provider string, attr authboss.Attributes) error {
	var authDataAttributes GoogleAuth
	if err := attr.Bind(&authDataAttributes, true); err != nil {
		return errors.Trace(err)
	}

	// Try and find the GoogleAuthData entry and update it, or create a new one
	var foundAuthData GoogleAuth
	authDataAttributes.Key = uid + provider
	err := store.database.
		Where(GoogleAuth{Oauth2Uid: uid, Oauth2Provider: provider}).
		Assign(authDataAttributes).
		FirstOrCreate(&foundAuthData).Error
	if err != nil {
		return errors.Trace(err)
	}
	return nil
}

// GetOAuth fetches the GoogleAuth record from the database
func (store *GoogleAuthStore) GetOAuth(uid string, provider string) (result interface{}, err error) {
	var foundAuthData GoogleAuth
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

func (store *GoogleAuthStore) migrate() {
	// Run auto migrations.
	migrate, migrateSet := os.LookupEnv(utils.Migrate)
	if migrateSet && migrate == "TRUE" {
		// Migrate schema soft changes
		store.database.AutoMigrate(
			&GoogleAuth{},
		)
	}
}
