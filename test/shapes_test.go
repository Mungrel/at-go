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
