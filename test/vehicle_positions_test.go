package test

import (
	"testing"

	. "github.com/Mungrel/at-go/client"
	"github.com/stretchr/testify/assert"
)

func TestVehiclePositions(t *testing.T) {
	mockATClient := newMockClient()
	assert := assert.New(t)

	vehiclePositions, err := mockATClient.VehiclePositions(nil, nil)
	assert.Nil(err)
	assert.NotNil(vehiclePositions)

	expected := []*VehiclePosition{
		{
			ID:                   "55bf6023-47bb-450f-e04f-bd47bc627fa0",
			IsDeleted:            false,
			TripID:               "440136659-20180702170310_v67.28",
			RouteID:              "06601-20180702170310_v67.28",
			StartTime:            "13:30:00",
			ScheduleRelationship: 0,
			VehicleID:            "5A57",
			Position: &Position{
				Latitude:  -36.91065,
				Longitude: 174.769567,
				Bearing:   "102",
			},
			Timestamp:       1531187743.075,
			OccupancyStatus: 0,
		},
		{
			ID:                   "870a00bb-13fa-8eb8-b951-6e821ca5d45b",
			IsDeleted:            false,
			TripID:               "473136639-20180702170310_v67.28",
			RouteID:              "02705-20180702170310_v67.28",
			StartTime:            "13:25:00",
			ScheduleRelationship: 0,
			VehicleID:            "2CB2",
			Position: &Position{
				Latitude:  -36.865117,
				Longitude: 174.76055,
				Bearing:   "320",
			},
			Timestamp:       1531187742.574,
			OccupancyStatus: 1,
		},
	}

	assert.Equal(expected, vehiclePositions)
}
