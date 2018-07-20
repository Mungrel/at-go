package test

import (
	"testing"

	. "github.com/Mungrel/at-go/client"
	"github.com/stretchr/testify/assert"
)

func TestStops(t *testing.T) {
	mockATClient := newMockClient()
	assert := assert.New(t)

	stops, err := mockATClient.Stops()
	assert.Nil(err)
	assert.NotNil(stops)

	expected := []*Stop{
		{
			ID:                 "0097-20180524131340_v66.89",
			Name:               "Papakura Train Station",
			Description:        "",
			Latitude:           -37.06429,
			Longitude:          174.94611,
			ZoneID:             "merged_90",
			URL:                "",
			Code:               "97",
			Street:             "",
			City:               "",
			Region:             "",
			PostCode:           "",
			Country:            "",
			LocationType:       0,
			ParentStation:      "11441-20180524131340_v66.89",
			Timezone:           "",
			WheelchairBoarding: "",
			Direction:          "",
			Position:           "",
			Geom:               "0101000020E6100000658D7A8846DE6540DCBA9BA73A8842C0",
		},
		{
			ID:                 "0098-20180524131340_v66.89",
			Name:               "Manurewa Train Station",
			Description:        "",
			Latitude:           -37.02327,
			Longitude:          174.89617,
			ZoneID:             "merged_90",
			URL:                "",
			Code:               "98",
			Street:             "",
			City:               "",
			Region:             "",
			PostCode:           "",
			Country:            "",
			LocationType:       0,
			ParentStation:      "",
			Timezone:           "",
			WheelchairBoarding: "",
			Direction:          "",
			Position:           "",
			Geom:               "0101000020E61000000135B56CADDC65402C7DE882FA8242C0",
		},
	}

	assert.Equal(expected, stops)
}

func TestStopsByID(t *testing.T) {
	mockATClient := newMockClient()
	assert := assert.New(t)

	stops, err := mockATClient.StopsByID(mockStopTimeStopID)
	assert.Nil(err)
	assert.NotNil(stops)

	expected := []*Stop{
		{
			ID:                 "0097-20180524131340_v66.89",
			Name:               "Papakura Train Station",
			Description:        "",
			Latitude:           -37.06429,
			Longitude:          174.94611,
			ZoneID:             "merged_90",
			URL:                "",
			Code:               "97",
			Street:             "",
			City:               "",
			Region:             "",
			PostCode:           "",
			Country:            "",
			LocationType:       0,
			ParentStation:      "",
			Timezone:           "",
			WheelchairBoarding: "",
			Direction:          "",
			Position:           "",
			Geom:               "0101000020E6100000658D7A8846DE6540DCBA9BA73A8842C0",
		},
	}

	assert.Equal(expected, stops)
}

func TestStopsByTripID(t *testing.T) {
	mockATClient := newMockClient()
	assert := assert.New(t)

	stops, err := mockATClient.StopsByTripID(mockStopTripID)
	assert.Nil(err)
	assert.NotNil(stops)

	expected := []*Stop{
		{
			ID:                 "2716-20180524131340_v66.89",
			Name:               "Papakura Train Station",
			Description:        "",
			Latitude:           -37.06496,
			Longitude:          174.9463,
			ZoneID:             "merged_90",
			URL:                "",
			Code:               "2716",
			Street:             "",
			Region:             "",
			City:               "",
			PostCode:           "",
			Country:            "",
			LocationType:       0,
			ParentStation:      "51012-20180524131340_v66.89",
			Timezone:           "",
			WheelchairBoarding: "",
			Direction:          "",
			Position:           "",
			Geom:               "0101000020E61000008E06F01648DE65402AC6F99B508842C0",
		},
		{
			ID:                 "2553-20180524131340_v66.89",
			Name:               "Railway St outside Countdown",
			Description:        "",
			Latitude:           -37.06365,
			Longitude:          174.94492,
			ZoneID:             "merged_90",
			URL:                "",
			Code:               "2553",
			Street:             "",
			Region:             "",
			City:               "",
			PostCode:           "",
			Country:            "",
			LocationType:       0,
			ParentStation:      "51614-20180524131340_v66.89",
			Timezone:           "",
			WheelchairBoarding: "",
			Direction:          "",
			Position:           "",
			Geom:               "0101000020E6100000C32ADEC83CDE6540F931E6AE258842C0",
		},
	}

	assert.Equal(expected, stops)
}

func TestStopsByLocation(t *testing.T) {
	mockATClient := newMockClient()
	assert := assert.New(t)

	stops, err := mockATClient.StopsByLocation(lat, lng, radius)
	assert.Nil(err)
	assert.NotNil(stops)

	expected := []*Stop{
		{
			ID:                 "0097-20180524131340_v66.89",
			Name:               "Papakura Train Station",
			Description:        "",
			Latitude:           -37.06429,
			Longitude:          174.94611,
			ZoneID:             "merged_90",
			URL:                "",
			Code:               "97",
			Street:             "",
			City:               "",
			Region:             "",
			PostCode:           "",
			Country:            "",
			LocationType:       0,
			ParentStation:      "",
			Timezone:           "",
			WheelchairBoarding: "",
			Direction:          "",
			Position:           "",
			Geom:               "0101000020E6100000658D7A8846DE6540DCBA9BA73A8842C0",
			StDistanceSphere:   0,
		},
		{
			ID:                 "51012-20180524131340_v66.89",
			Name:               "Papakura",
			Description:        "",
			Latitude:           -37.06462,
			Longitude:          174.94572,
			ZoneID:             "merged_90",
			URL:                "",
			Code:               "51012",
			Street:             "",
			City:               "",
			Region:             "",
			PostCode:           "",
			Country:            "",
			LocationType:       1,
			ParentStation:      "",
			Timezone:           "",
			WheelchairBoarding: "",
			Direction:          "",
			Position:           "",
			Geom:               "0101000020E61000008AE5965643DE65407155D977458842C0",
			StDistanceSphere:   50.437442361,
		},
	}

	assert.Equal(expected, stops)
}

func TestStopsSearch(t *testing.T) {
	mockATClient := newMockClient()
	assert := assert.New(t)

	stops, err := mockATClient.StopsSearch(mockSearchTerm)
	assert.Nil(err)
	assert.NotNil(stops)

	expected := []*Stop{
		{
			ID:                 "51012-20180524131340_v66.89",
			Name:               "Papakura",
			Description:        "",
			Latitude:           -37.06462,
			Longitude:          174.94572,
			ZoneID:             "merged_90",
			URL:                "",
			Code:               "51012",
			Street:             "",
			City:               "",
			Region:             "",
			PostCode:           "",
			Country:            "",
			LocationType:       1,
			ParentStation:      "",
			Timezone:           "",
			WheelchairBoarding: "",
			Direction:          "",
			Position:           "",
			Geom:               "0101000020E61000008AE5965643DE65407155D977458842C0",
		},
		{
			ID:                 "2716-20180524131340_v66.89",
			Name:               "Papakura Train Station",
			Description:        "",
			Latitude:           -37.06496,
			Longitude:          174.9463,
			ZoneID:             "merged_90",
			URL:                "",
			Code:               "2716",
			Street:             "",
			City:               "",
			Region:             "",
			PostCode:           "",
			Country:            "",
			LocationType:       0,
			ParentStation:      "51012-20180524131340_v66.89",
			Timezone:           "",
			WheelchairBoarding: "",
			Direction:          "",
			Position:           "",
			Geom:               "0101000020E61000008E06F01648DE65402AC6F99B508842C0",
		},
	}

	assert.Equal(expected, stops)
}

func TestStopsByCode(t *testing.T) {
	mockATClient := newMockClient()
	assert := assert.New(t)

	stops, err := mockATClient.StopsByCode(mockStopCode)
	assert.Nil(err)
	assert.NotNil(stops)

	expected := []*Stop{
		{
			ID:                 "2716-20180702170310_v67.28",
			Name:               "Papakura Train Station",
			Description:        "",
			Latitude:           -37.06496,
			Longitude:          174.9463,
			ZoneID:             "merged_91",
			URL:                "",
			Code:               "2716",
			Street:             "",
			City:               "",
			Region:             "",
			PostCode:           "",
			Country:            "",
			LocationType:       0,
			ParentStation:      "51012-20180702170310_v67.28",
			Timezone:           "",
			WheelchairBoarding: "",
			Direction:          "",
			Position:           "",
			Geom:               "0101000020E61000008E06F01648DE65402AC6F99B508842C0",
		},
	}

	assert.Equal(expected, stops)
}

func TestStopInfoByCode(t *testing.T) {
	mockATClient := newMockClient()
	assert := assert.New(t)

	stopInfos, err := mockATClient.StopInfoByCode(mockStopCode)
	assert.Nil(err)
	assert.NotNil(stopInfos)

	expected := []*StopInfo{
		{
			TripID:         "2365121452-20180524131340_v66.89",
			DepartureTime:  "07:50:00",
			TripShortName:  "",
			TripHeadsign:   "Manukau",
			RouteLongName:  "Papakura Interchange To Manukau Bus Station Via Porchester",
			RouteShortName: "365",
			StopSequence:   1,
			PickupType:     0,
			DropOffType:    0,
		},
		{
			TripID:         "1064099009-20180524131340_v66.89",
			DepartureTime:  "07:50:00",
			TripShortName:  "",
			TripHeadsign:   "Otahuhu",
			RouteLongName:  "Papakura Interchange To Otahuhu Station Via Great South Rd",
			RouteShortName: "33",
			StopSequence:   1,
			PickupType:     0,
			DropOffType:    0,
		},
	}

	assert.Equal(expected, stopInfos)
}
