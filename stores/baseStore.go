package stores

import (
	"github.com/jinzhu/gorm"
)

// BaseStore represents a Store abstraction collection specific stores.
type BaseStore struct {
	Database *gorm.DB
}
