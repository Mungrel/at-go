package test

import (
	"testing"

	. "github.com/Mungrel/at-go"
	"github.com/stretchr/testify/assert"
)

func TestVersions(t *testing.T) {
	mockATClient := newMockClient()
	assert := assert.New(t)

	versions, err := mockATClient.Versions()
	assert.Nil(err)
	assert.NotNil(versions)

	expected := []*Version{
		{
			Version:   "20180524131340_v66.89",
			StartDate: NewTimestamp("2018-05-17T00:00:00.000Z"),
			EndDate:   NewTimestamp("2018-07-07T00:00:00.000Z"),
		},
		{
			Version:   "20180605110613_v67.1",
			StartDate: NewTimestamp("2018-07-08T00:00:00.000Z"),
			EndDate:   NewTimestamp("2018-09-30T00:00:00.000Z"),
		},
	}

	assert.Equal(expected, versions)
}
