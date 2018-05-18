package models

import (
	"time"
)

// UserSession :
type UserSession struct {
	SessionID     string    `gorm:"primary_key;not null" json:"session_id"`
	UserID        uint      `gorm:"type:int(11) unsigned;not null" json:"user_id"`
	SignedInDate  time.Time `gorm:"type:datetime;column:signed_in;not null" json:"signed_in"`
	RemoteAddress string    `gorm:"column:remote_addr" json:"remote_addr"`
}
