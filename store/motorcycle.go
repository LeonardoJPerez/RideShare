package store

import (
	"github.com/RideShare-Server/models"
)

type MotorcycleStore struct {
	BaseStore
}

func NewChannelStore() *MotorcycleStore {
	m := new(models.Motorcycle)

	store := new(MotorcycleStore)
	store.CollectionName = m.CollectionName()

	return store
}

func (store *MotorcycleStore) GetMotorcycle(id string) (*models.Motorcycle, error) {
	// whereClause := model.Channel{
	// 	RemoteID: remoteID,
	// 	Provider: provider,
	// }

	// var channel model.Channel
	// err := store.Database.Where(whereClause).First(&channel).Error
	// if err != nil {
	// 	return nil, errors.Trace(err)
	// }

	return nil, nil
}
