package test

import (
	"testing"

	. "github.com/Mungrel/at-go/client"
	"github.com/stretchr/testify/assert"
)

func TestStopTimesByID(t *testing.T) {
	mockATClient := newMockClient()
	assert := assert.New(t)

	stopTimes, err := mockATClient.StopTimesByID(mockStopTimeStopID)
	assert.Nil(err)
	assert.NotNil(stopTimes)

	expected := []*StopTime{
		{
			TripID:                 "23100072400-20180524131340_v66.89",
			ArrivalTime:            "14:25:59",
			DepartureTime:          "14:25:59",
			StopID:                 "0097-20180524131340_v66.89",
			StopSequence:           16,
			StopHeadsign:           "",
			PickupType:             1,
			DropoffType:            0,
			ShapeDistanceTravelled: 0.0,
			Timepoint:              "",
			ArrivalTimeSeconds:     51959,
			DepartureTimeSeconds:   51959,
		},
		{
			TripID:                 "23100071244-20180524131340_v66.89",
			ArrivalTime:            "17:37:00",
			DepartureTime:          "17:37:00",
			StopID:                 "0097-20180524131340_v66.89",
			StopSequence:           1,
			StopHeadsign:           "",
			PickupType:             0,
			DropoffType:            0,
			ShapeDistanceTravelled: 0.0,
			Timepoint:              "",
			ArrivalTimeSeconds:     63420,
			DepartureTimeSeconds:   63420,
		},
	}

	assert.Equal(expected, stopTimes)
}
