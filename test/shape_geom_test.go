package test

import (
	"testing"

	. "github.com/Mungrel/at-go/client"
	"github.com/stretchr/testify/assert"
)

func TestShapeGeomByID(t *testing.T) {
	mockATClient := newMockClient()
	assert := assert.New(t)

	shapeGeom, err := mockATClient.ShapeGeometryByID(mockShapeGeomID)
	assert.Nil(err)
	assert.NotNil(shapeGeom)

	expected := []*ShapeGeom{
		{
			ShapeID: "813-20180524131340_v66.89",
			Geom:    "0102000020E61000006F0400004EB9C2BB5CD56540B9196EC0E74F42C04EB9C2BB",
		},
	}

	assert.Equal(expected, shapeGeom)
}
