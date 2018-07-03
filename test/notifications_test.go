package test

import (
	"testing"

	. "github.com/Mungrel/at-go/client"
	"github.com/stretchr/testify/assert"
)

func TestNotifications(t *testing.T) {
	mockATClient := newMockClient()
	assert := assert.New(t)

	notifications, err := mockATClient.Notifications()
	assert.Nil(err)
	assert.NotNil(notifications)

	expected := &Notifications{
		Data: []*Notification{
			{
				ID:              "1",
				OrderingDate:    NewTimestamp("2018-07-03 10:48:26"),
				Type:            "MOVED_STOP",
				Title:           "Stop 7004 - Sturdee St near Pakenham St has been moved",
				StopID:          7004,
				ExistingAddress: "Sturdee St near Pakenham St",
				Description:     "The following services have been affected: 76X,802X,820,822,834,837,839,858,85X,863X,866X,86X,873X,874X,875,877X,879,87X,900X,920,922,945X,952,956",
			},
			{
				ID:                 "219060",
				OrderingDate:       NewTimestamp("2018-07-03 09:00:00"),
				Type:               "HIGHWAY",
				Title:              "Highway maintenance",
				Description:        "Delays may be experienced by the public due to the ramp closure. This road work closure is dependent on various factors, including weather, and may be postponed or cancelled at late notice",
				Impact:             "Road Closed",
				StartDate:          NewTimestamp("2018-07-03 09:00:00"),
				EndDate:            NewTimestamp("2018-07-03 18:00:00"),
				ExpectedResolution: "04/07/2018 06:00",
				Location: &Location{
					Latitude:  -36.96320874052837,
					Longitude: 174.86758125862653,
					Address:   "12 Natalie Place, Otara\nAuckland 2023",
				},
			},
		},
		Time: NewTimestamp("2018-07-03 10:48:36"),
	}

	assert.Equal(expected, notifications)
}