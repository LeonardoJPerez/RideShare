package stores

import (
	"github.com/RideShare-Server/db"
	"github.com/RideShare-Server/models"

	"github.com/jinzhu/gorm"
	"github.com/juju/errors"
)

// MotorcycleStore is responsible for persisting and querying Motorcycle objects.
type MotorcycleStore struct {
	BaseStore
}

// NewMotorcycleStore :
func NewMotorcycleStore(database *gorm.DB) *MotorcycleStore {
	if database == nil {
		panic("Database must not be nil")
	}

	s := new(MotorcycleStore)
	s.Database = database

	return s
}

// GetByID :
func (store *MotorcycleStore) GetByID(bikeID uint) (*models.Motorcycle, error) {
	m := &models.Motorcycle{}
	err := store.Database.
		First(m, bikeID).
		Error

	if err != nil {
		return nil, errors.Trace(db.CheckNotFoundErr(err))
	}

	return m, nil
}

// GetByUser :
func (store *MotorcycleStore) GetByUser(userID uint) ([]*models.Motorcycle, error) {
	garage := []*models.Motorcycle{}
	err := store.Database.
		Where("user_id = ?", userID).
		Find(&garage).
		Error
	if err != nil {
		return nil, errors.Trace(db.CheckNotFoundErr(err))
	}

	return garage, nil
}

// Insert :
func (store *MotorcycleStore) Insert(g *models.Motorcycle) (*models.Motorcycle, error) {
	err := store.Database.Create(g).Error
	if err != nil {
		return nil, err
	}

	return g, nil
}

// Remove :
func (store *MotorcycleStore) Remove(id uint) (bool, error) {
	m := new(models.Motorcycle)
	m.ID = id

	err := store.Database.Delete(m).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

// SetPreferredBike :
func (store *MotorcycleStore) SetPreferredBike(userID, bikeID uint) (bool, error) {
	bike := new(models.Motorcycle)
	bike.ID = userID

	pros := map[string]interface{}{
		"preferred_bike_id": bikeID}
	err := store.Database.
		Model(bike).
		Select("preferred_bike_id").
		Updates(pros).Error
	if err != nil {
		return false, err
	}

	return true, nil
}
