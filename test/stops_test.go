package test

import (
	"testing"

	. "github.com/Mungrel/at-go"
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