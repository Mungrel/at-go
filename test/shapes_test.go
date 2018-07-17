package test

import (
	"testing"

	. "github.com/Mungrel/at-go/client"
	"github.com/stretchr/testify/assert"
)

func TestShapesByID(t *testing.T) {
	mockATClient := newMockClient()
	assert := assert.New(t)

	shapes, err := mockATClient.ShapesByID(mockShapeID)
	assert.Nil(err)
	assert.NotNil(shapes)

	expected := []*Shape{
		{
			ID:                "81-20180605110613_v67.1",
			PointLat:          -36.72234,
			PointLong:         174.7125,
			PointSequence:     1,
			DistanceTravelled: float64(0),
		},
		{
			ID:                "81-20180605110613_v67.1",
			PointLat:          -36.72234,
			PointLong:         174.7125,
			PointSequence:     2,
			DistanceTravelled: float64(0),
		},
	}

	assert.Equal(expected, shapes)
}

func TestShapesByTrip(t *testing.T) {
	mockATClient := newMockClient()
	assert := assert.New(t)

	shapes, err := mockATClient.ShapesByTrip(mockShapeTripID)
	assert.Nil(err)
	assert.NotNil(shapes)

	expected := []*Shape{
		{
			ID:                "854-20180524131340_v66.89",
			PointLat:          -37.06496,
			PointLong:         174.9463,
			PointSequence:     1,
			DistanceTravelled: 0.0,
		},
		{
			ID:                "854-20180524131340_v66.89",
			PointLat:          -37.06496,
			PointLong:         174.9463,
			PointSequence:     2,
			DistanceTravelled: 0.0,
		},
	}

	assert.Equal(expected, shapes)
}
