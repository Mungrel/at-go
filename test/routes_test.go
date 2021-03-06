package test

import (
	"testing"

	. "github.com/Mungrel/at-go/client"
	"github.com/stretchr/testify/assert"
)

const (
	lat    = -37.06429
	lng    = 174.94611
	radius = 500.0
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

func TestRoutesByLocation(t *testing.T) {
	mockATClient := newMockClient()
	assert := assert.New(t)

	routes, err := mockATClient.RoutesByLocation(lat, lng, radius)
	assert.Nil(err)
	assert.NotNil(routes)

	expected := []*Route{
		{
			ID:               "20439-20180524131340_v66.89",
			AgencyID:         "AM",
			ShortName:        "STH",
			LongName:         "Papakura Train Station to Britomart Train Station",
			Description:      "",
			Type:             2,
			URL:              "",
			Color:            "",
			TextColor:        "",
			StDistanceSphere: 127.33146307,
		},
		{
			ID:               "20328-20180524131340_v66.89",
			AgencyID:         "AM",
			ShortName:        "STH",
			LongName:         "Britomart Train Station to Papakura Train Station",
			Description:      "",
			Type:             2,
			URL:              "",
			Color:            "",
			TextColor:        "",
			StDistanceSphere: 127.33146307,
		},
	}

	assert.Equal(expected, routes)
}

func TestRoutesByLongName(t *testing.T) {
	mockATClient := newMockClient()
	assert := assert.New(t)

	routes, err := mockATClient.RoutesByLongName(mockRoutesLongName)
	assert.Nil(err)
	assert.NotNil(routes)

	expected := []*Route{
		{
			ID:          "20439-20180524131340_v66.89",
			AgencyID:    "AM",
			ShortName:   "STH",
			LongName:    "Papakura Train Station to Britomart Train Station",
			Description: "",
			Type:        2,
			URL:         "",
			Color:       "",
			TextColor:   "",
		},
		{
			ID:          "20439-20180605110613_v67.1",
			AgencyID:    "AM",
			ShortName:   "STH",
			LongName:    "Papakura Train Station to Britomart Train Station",
			Description: "",
			Type:        2,
			URL:         "",
			Color:       "",
			TextColor:   "",
		},
	}

	assert.Equal(expected, routes)
}

func TestRoutesByShortName(t *testing.T) {
	mockATClient := newMockClient()
	assert := assert.New(t)

	routes, err := mockATClient.RoutesByShortName(mockRoutesShortName)
	assert.Nil(err)
	assert.NotNil(routes)

	expected := []*Route{
		{
			ID:          "20439-20180524131340_v66.89",
			AgencyID:    "AM",
			ShortName:   "STH",
			LongName:    "Papakura Train Station to Britomart Train Station",
			Description: "",
			Type:        2,
			URL:         "",
			Color:       "",
			TextColor:   "",
		},
		{
			ID:          "20439-20180605110613_v67.1",
			AgencyID:    "AM",
			ShortName:   "STH",
			LongName:    "Papakura Train Station to Britomart Train Station",
			Description: "",
			Type:        2,
			URL:         "",
			Color:       "",
			TextColor:   "",
		},
	}

	assert.Equal(expected, routes)
}

func TestRoutesByStopID(t *testing.T) {
	mockATClient := newMockClient()
	assert := assert.New(t)

	routes, err := mockATClient.RoutesByStopID(mockStopID)
	assert.Nil(err)
	assert.NotNil(routes)

	expected := []*Route{
		{
			ID:          "20419-20180524131340_v66.89",
			AgencyID:    "AM",
			ShortName:   "STH",
			LongName:    "Papakura Train Station to Britomart Train Station",
			Description: "",
			Type:        2,
			URL:         "",
			Color:       "",
			TextColor:   "",
		},
		{
			ID:          "20437-20180524131340_v66.89",
			AgencyID:    "AM",
			ShortName:   "STH",
			LongName:    "Papakura Train Station to Britomart Train Station",
			Description: "",
			Type:        2,
			URL:         "",
			Color:       "",
			TextColor:   "",
		},
	}

	assert.Equal(expected, routes)
}

func TestRoutesByID(t *testing.T) {
	mockATClient := newMockClient()
	assert := assert.New(t)

	routes, err := mockATClient.RoutesByID(mockRouteID)
	assert.Nil(err)
	assert.NotNil(routes)

	expected := []*Route{
		{
			ID:          "20439-20180524131340_v66.89",
			AgencyID:    "AM",
			ShortName:   "STH",
			LongName:    "Papakura Train Station to Britomart Train Station",
			Description: "",
			Type:        2,
			URL:         "",
			Color:       "",
			TextColor:   "",
		},
	}

	assert.Equal(expected, routes)
}

func TestRoutesSearch(t *testing.T) {
	mockATClient := newMockClient()
	assert := assert.New(t)

	routes, err := mockATClient.RoutesSearch(mockSearchTerm)
	assert.Nil(err)
	assert.NotNil(routes)

	expected := []*Route{
		{
			ID:          "37101-20180702170310_v67.28",
			AgencyID:    "RTH",
			ShortName:   "371",
			LongName:    "Takanini Station To Papakura Interchange",
			Description: "",
			Type:        3,
			URL:         "",
			Color:       "",
			TextColor:   "",
		},
		{
			ID:          "05555-20180524131340_v66.89",
			AgencyID:    "RTH",
			ShortName:   "055",
			LongName:    "Papakura To McAuley High",
			Description: "",
			Type:        3,
			URL:         "",
			Color:       "",
			TextColor:   "",
		},
	}

	assert.Equal(expected, routes)
}
