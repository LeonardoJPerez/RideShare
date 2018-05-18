package models

import "time"

// GoogleAuth represents the data we get from Google in the auth process
type GoogleAuth struct {
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
