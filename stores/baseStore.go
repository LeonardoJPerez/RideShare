package stores

import (
	"github.com/RideShare-Server/db"
	"github.com/jinzhu/gorm"
	"github.com/juju/errors"
)

var (
	// Base store.
	Base *BaseStore
	// Motorcycles :
	Motorcycles *motorcycleStore
	// Rides :
	Rides *rideStore
)

// Init :
func Init(db *gorm.DB) {
	Rides = newrideStore(db)
	Motorcycles = newMotorcycleStore(db)
	Base = &Motorcycles.BaseStore
}

// BaseStore represents a Store abstraction collection specific stores.
type BaseStore struct {
	database *gorm.DB
}

// Insert inserts a Resource.
func (store *BaseStore) Insert(m interface{}) (interface{}, error) {
	err := store.database.Create(m).Error
	if err != nil {
		return nil, err
	}

	return m, nil
}

// Get fetches a Resource.
func (store *BaseStore) Get(id uint, m interface{}) (interface{}, error) {
	err := store.database.
		First(m, id).
		Error

	if err != nil {
		return nil, errors.Trace(db.CheckNotFoundErr(err))
	}

	return m, nil
}

// Update updates a Resource.
func (store *BaseStore) Update(kvp, m interface{}) (interface{}, error) {
	err := store.database.
		Model(m).
		Updates(kvp).Error
	if err != nil {
		return nil, err
	}

	return m, nil
}

// Remove removes a Resource.
func (store *BaseStore) Remove(id uint, m interface{}) (bool, error) {
	err := store.database.Where("id = ?", id).Delete(m).Error
	if err != nil {
		return false, err
	}

	return true, nil
}
