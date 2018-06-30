package test

import (
	"testing"

	. "github.com/Mungrel/at-go/client"
	"github.com/stretchr/testify/assert"
)

func TestParkingLocations(t *testing.T) {
	mockATClient := newMockClient()
	assert := assert.New(t)

	parkingLocations, err := mockATClient.ParkingLocations()
	assert.Nil(err)
	assert.NotNil(parkingLocations)

	expected := []*ParkingLocation{
		{
			ID:             "1",
			Name:           "P120 parallel park without kerb cutting.",
			Address:        "16 Birkenhead Avenue Birkenhead Birkenhead",
			MobilitySpaces: 1,
			Type:           "mobility",
			Latitude:       -36.811614,
			Longitude:      174.726024,
		},
		{
			ID:             "2",
			Name:           "Not time limted angle parking in parking area at end of Anzac Road.Operates at all times.Drop kerb access to rear of vehicles.",
			Address:        " Anzac Street Browns Bay Browns Bay",
			MobilitySpaces: 2,
			Type:           "mobility",
			Latitude:       -36.716683,
			Longitude:      174.749114,
		},
	}

	assert.Equal(expected, parkingLocations)
}
