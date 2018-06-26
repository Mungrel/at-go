package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMetadata(t *testing.T) {
	mockATClient := newMockClient()
	assert := assert.New(t)

	resp, err := mockATClient.Metadata(nil)
	assert.Nil(err)
	assert.NotNil(resp)

	expected := "534472888"

	assert.Equal(&expected, resp)
}
