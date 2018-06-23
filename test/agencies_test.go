package test

import (
	"testing"

	. "github.com/Mungrel/at-go"
	"github.com/stretchr/testify/assert"
)

func TestAgencies(t *testing.T) {
	mockATClient := newMockClient()
	assert := assert.New(t)

	agencies, err := mockATClient.Agencies()
	assert.Nil(err)
	assert.NotNil(agencies)

	expected := []*Agency{
		&Agency{
			ID:       "NZBGW",
			Name:     "Go West",
			URL:      "http://www.aucklandtransport.govt.nz",
			Timezone: "Pacific/Auckland",
			Language: "en",
			Phone:    "(09)355-3553",
			FareURL:  "",
		},
		&Agency{
			ID:       "SLPH",
			Name:     "SeaLink Pine Harbour",
			URL:      "http://www.aucklandtransport.govt.nz",
			Timezone: "Pacific/Auckland",
			Language: "en",
			Phone:    "(09)355-3553",
			FareURL:  "",
		},
	}

	assert.Equal(expected, agencies)
}
