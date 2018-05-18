package stores

import (
	"github.com/RideShare-Server/db"
	"github.com/RideShare-Server/models"

	"github.com/jinzhu/gorm"
	"github.com/juju/errors"
)

// MotorcycleStore is responsible for persisting and querying Motorcycle objects.
type motorcycleStore struct {
	BaseStore
}

// NewMotorcycleStore :
func newMotorcycleStore(database *gorm.DB) *motorcycleStore {
	if database == nil {
		panic("Database must not be nil")
	}

	s := new(motorcycleStore)
	s.database = database

	return s
}

// GetByUser :
func (store *motorcycleStore) GetByUser(userID uint) ([]*models.Motorcycle, error) {
	garage := []*models.Motorcycle{}
	err := store.database.
		Where("user_id = ?", userID).
		Find(&garage).
		Error
	if err != nil {
		return nil, errors.Trace(db.CheckNotFoundErr(err))
	}

	return garage, nil
}

// SetPreferredBike :
func (store *motorcycleStore) SetPreferredBike(userID, bikeID uint) (bool, error) {
	bike := new(models.Motorcycle)
	bike.ID = userID

	pros := map[string]interface{}{
		"preferred_bike_id": bikeID}
	err := store.database.
		Model(bike).
		Select("preferred_bike_id").
		Updates(pros).Error
	if err != nil {
		return false, err
	}

	return true, nil
}
