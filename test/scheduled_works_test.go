package test

import (
	"testing"

	. "github.com/Mungrel/at-go/client"
	"github.com/stretchr/testify/assert"
)

func TestScheduledWorks(t *testing.T) {
	mockATClient := newMockClient()
	assert := assert.New(t)

	scheduledWorks, err := mockATClient.ScheduledWorks()
	assert.Nil(err)
	assert.NotNil(scheduledWorks)

	expected := []*ScheduledWorks{
		{
			ID:                "109847",
			Code:              "E109847",
			Type:              "Non Excavation",
			Name:              "Non Excavation",
			Address:           "",
			Suburb:            "Auckland",
			Region:            "Auckland Central",
			Description:       "Other (Specify Below). To install and/or remove cross street banners or Flags to poles Trucks used to protect cherry picker when working: ? Level 1 - any speed level 1 truck with arrow board  ? Level 2 -under 70kph needs 1 attenuator  ? Level 2 - 70kph + needs 2 attenuator  Any State Highway (Level 1 or 2) needs NZTA separate approval  Flags As listed in the following sheets & maps attached.",
			StartDate:         NewTimestamp("2013-06-01T00:00:00Z"),
			EndDate:           NewTimestamp("2015-06-30T00:00:00Z"),
			LastUpdated:       NewTimestamp("2017-03-17T10:48:23Z"),
			ContractorCompany: "NZ Traffic Ltd",
			Latitude:          -36.84845969999998,
			Longitude:         174.76333150000005,
			Geometry: &Geometry{
				Type: "point",
				Coordinates: [][]float64{
					{
						174.76333150000005,
						-36.84845969999998,
					},
				},
				Encoded: false,
			},
		},
		{
			ID:                "185304",
			Code:              "583225",
			Type:              "Excavation",
			Name:              "Excavation",
			Address:           "229 Richardson Road, Mount Roskill South, Auckland on Carriageway and Footpath and Berm",
			Suburb:            "Mount Roskill South",
			Region:            "Auckland Central",
			Description:       "Minor Earthworks/Filling. Richardson Road",
			StartDate:         NewTimestamp("2014-06-17T00:00:00Z"),
			EndDate:           NewTimestamp("2017-10-31T00:00:00Z"),
			LastUpdated:       NewTimestamp("2017-04-19T13:37:54Z"),
			ContractorCompany: "FLETCHER CONSTRUCTION LIMITED",
			Latitude:          -36.901927196999964,
			Longitude:         174.7196914980001,
			Geometry: &Geometry{
				Type: "polygon",
				Coordinates: [][]float64{
					{
						-36.90203020499996,
						174.71956268500003,
					},
					{
						-36.90202745199997,
						174.71958442200003,
					},
					{
						-36.902022012999964,
						174.71960535200003,
					},
					{
						-36.90201402399998,
						174.71962496100002,
					},
				},
				Encoded: false,
			},
		},
	}

	assert.Equal(expected, scheduledWorks)
}
