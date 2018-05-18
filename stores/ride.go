package stores

import (
	"time"

	"github.com/RideShare-Server/db"
	"github.com/RideShare-Server/log"
	"github.com/RideShare-Server/models"
	"github.com/RideShare-Server/utils"
	"github.com/jinzhu/gorm"
	"github.com/juju/errors"
)

type rideStore struct {
	BaseStore
}

func newrideStore(database *gorm.DB) *rideStore {
	if database == nil {
		panic("Database must not be nil")
	}

	s := new(rideStore)
	s.database = database

	return s
}

func (store *rideStore) GetStartingBy(from, to string) ([]*models.Ride, error) {
	events := []*models.Ride{}
	if from == "" {
		return events, errors.New("'from' datestamp cannot be empty")
	}

	// Get starting date.
	fromDate, err := utils.ParseEpoch(from)
	if err != nil {
		return events, err
	}

	// Get upper limit.
	var toDate *time.Time
	if to != "" {
		toDate, err = utils.ParseEpoch(to)
		if err != nil {
			return events, err
		}
	}

	query := store.database.Where("when > ", from)
	if toDate != nil {
		query = query.Where("when BETWEEN ? AND ?", fromDate, toDate)
	}

	err = query.Find(&events).Error
	if err != nil {
		err := errors.Trace(db.CheckNotFoundErr(err))
		log.Error(log.StoreLayerTopic, err)
		return events, err
	}

	return events, nil
}

func (store *rideStore) GetByOwnerID(ownerID uint)                   {}
func (store *rideStore) GetByTags(tags []string)                     {}
func (store *rideStore) GetByLocation(startLocation *models.Address) {}
