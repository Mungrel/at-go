package test

import (
	"testing"

	. "github.com/Mungrel/at-go/client"
	"github.com/stretchr/testify/assert"
)

func TestCustomerServiceCentres(t *testing.T) {
	mockATClient := newMockClient()
	assert := assert.New(t)

	csc, err := mockATClient.CustomerServiceCentres()
	assert.Nil(err)
	assert.NotNil(csc)

	expected := []*CustomerServiceCentre{
		{
			ID:          "1",
			Name:        "Newmarket Train Station",
			Address:     "48A Remuera Rd",
			Suburb:      "Newmarket",
			Phone:       "09 366 6400",
			OpenHours:   "Mon- Fri: 7am-6pm\nSat: 8am-4.30pm\nClosed Sunday",
			HopServices: true,
			Latitude:    -36.869167,
			Longitude:   -36.869167,
		},
		{
			ID:          "2",
			Name:        "AUT City Campus",
			Address:     "55 Building WA 4/11, Hikuwai Plaza, Wellesley St East",
			Suburb:      "CBD",
			Phone:       "09 366 6400",
			OpenHours:   "Mon – Fri: 8.30am-5pm\nClosed Saturday and Sunday",
			HopServices: true,
			Latitude:    -36.852986,
			Longitude:   -36.852986,
		},
	}

	assert.Equal(expected, csc)
}
