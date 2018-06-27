package test

import (
	"testing"

	. "github.com/Mungrel/at-go"
	"github.com/stretchr/testify/assert"
)

func TestCalendars(t *testing.T) {
	mockATClient := newMockClient()
	assert := assert.New(t)

	calendars, err := mockATClient.Calendars()
	assert.Nil(err)
	assert.NotNil(calendars)

	expected := []*Calendar{
		{
			ServiceID: "18033099921-20180524131340_v66.89",
			Monday:    0,
			Tuesday:   0,
			Wednesday: 0,
			Thursday:  0,
			Friday:    0,
			Saturday:  0,
			Sunday:    0,
			StartDate: NewTimestamp("2018-05-17T00:00:00.000Z"),
			EndDate:   NewTimestamp("2018-07-07T00:00:00.000Z"),
		},
		{
			ServiceID: "12881116336-20180605110613_v67.1",
			Monday:    1,
			Tuesday:   1,
			Wednesday: 1,
			Thursday:  1,
			Friday:    1,
			Saturday:  0,
			Sunday:    0,
			StartDate: NewTimestamp("2018-07-08T00:00:00.000Z"),
			EndDate:   NewTimestamp("2018-09-30T00:00:00.000Z"),
		},
	}

	assert.Equal(expected, calendars)
}

func TestCalendarByService(t *testing.T) {
	mockATClient := newMockClient()
	assert := assert.New(t)

	calendars, err := mockATClient.CalendarByService(mockServiceID)
	assert.Nil(err)
	assert.NotNil(calendars)

	expected := []*Calendar{
		{
			ServiceID: "18033099921-20180524131340_v66.89",
			Monday:    0,
			Tuesday:   0,
			Wednesday: 0,
			Thursday:  0,
			Friday:    0,
			Saturday:  0,
			Sunday:    0,
			StartDate: NewTimestamp("2018-05-17T00:00:00.000Z"),
			EndDate:   NewTimestamp("2018-07-07T00:00:00.000Z"),
		},
	}

	assert.Equal(expected, calendars)
}
