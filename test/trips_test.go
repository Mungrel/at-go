package test

import (
	"testing"

	. "github.com/Mungrel/at-go"
	"github.com/stretchr/testify/assert"
)

func TestTrips(t *testing.T) {
	mockATClient := newMockClient()
	assert := assert.New(t)

	trips, err := mockATClient.Trips()
	assert.Nil(err)
	assert.NotNil(trips)

	expected := []*Trip{
		{
			RouteID:       "98302-20180524131340_v66.89",
			ServiceID:     "1983053937-20180524131340_v66.89",
			TripID:        "1983053937-20180524131340_v66.89",
			TripHeadsign:  "Gulf Harbour",
			DirectionID:   1,
			BlockID:       "",
			ShapeID:       "813-20180524131340_v66.89",
			TripShortName: "",
			TripType:      "",
		},
		{
			RouteID:       "88103-20180605110613_v67.1",
			ServiceID:     "12881116336-20180605110613_v67.1",
			TripID:        "12881116336-20180605110613_v67.1",
			TripHeadsign:  "Newmarket",
			DirectionID:   0,
			BlockID:       "",
			ShapeID:       "81-20180605110613_v67.1",
			TripShortName: "",
			TripType:      "",
		},
	}

	assert.Equal(expected, trips)
}

func TestTripsByServiceID(t *testing.T) {
	mockATClient := newMockClient()
	assert := assert.New(t)

	trips, err := mockATClient.TripsByID(mockTripID)
	assert.Nil(err)
	assert.NotNil(trips)

	expected := []*Trip{
		{
			RouteID:       "98302-20180524131340_v66.89",
			ServiceID:     "1983053937-20180524131340_v66.89",
			TripID:        "1983053937-20180524131340_v66.89",
			TripHeadsign:  "Gulf Harbour",
			DirectionID:   1,
			BlockID:       "",
			ShapeID:       "813-20180524131340_v66.89",
			TripShortName: "",
			TripType:      "",
		},
	}

	assert.Equal(expected, trips)
}
