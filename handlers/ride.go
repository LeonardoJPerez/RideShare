package handlers

// RideHandler :
type RideHandler struct {
	Base
}

// NewRideHandler :
func NewRideHandler() *RideHandler {
	c := new(RideHandler)

	return c
}

/*
Methods
- GetRidesByLocation(startLocation *Address)
- GetRidesStartingBy(from, to string)
- GetRidesByOwnerID(ownerID uint)
- GetRidesByTags(tags []string)
*/
