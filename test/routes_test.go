package test

import (
	"testing"

	. "github.com/Mungrel/at-go"
	"github.com/stretchr/testify/assert"
)

func TestRoutes(t *testing.T) {
	mockATClient := newMockClient()
	assert := assert.New(t)

	routes, err := mockATClient.Routes()
	assert.Nil(err)
	assert.NotNil(routes)

	expected := []*Route{
		{
			ID:          "65001-20180605110613_v67.1",
			AgencyID:    "NZB",
			ShortName:   "650",
			LongName:    "Pt Chevalier Shops To Glen Innes Via Greenlane",
			Description: "",
			Type:        3,
			URL:         "",
			Color:       "",
			TextColor:   "",
		},
		{
			ID:          "35802-20180524131340_v66.89",
			AgencyID:    "HE",
			ShortName:   "358",
			LongName:    "Onehunga To Pakuranga Plaza Via Penrose",
			Description: "",
			Type:        3,
			URL:         "",
			Color:       "",
			TextColor:   "",
		},
	}

	assert.Equal(expected, routes)
}
