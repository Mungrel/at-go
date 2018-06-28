package test

import (
	"testing"

	. "github.com/Mungrel/at-go/client"
	"github.com/stretchr/testify/assert"
)

func TestCalendarDates(t *testing.T) {
	mockATClient := newMockClient()
	assert := assert.New(t)

	calendarDates, err := mockATClient.CalendarDates()
	assert.Nil(err)
	assert.NotNil(calendarDates)

	expected := []*CalendarDate{
		{
			ServiceID:     "18033099921-20180524131340_v66.89",
			Date:          NewTimestamp("2018-05-17T00:00:00.000Z"),
			ExceptionType: 1,
		},
		{
			ServiceID:     "18033099921-20180524131340_v66.89",
			Date:          NewTimestamp("2018-05-20T00:00:00.000Z"),
			ExceptionType: 1,
		},
	}

	assert.Equal(expected, calendarDates)
}

func TestCalendarDateByService(t *testing.T) {
	mockATClient := newMockClient()
	assert := assert.New(t)

	calendarDates, err := mockATClient.CalendarDateByService(mockServiceID)
	assert.Nil(err)
	assert.NotNil(calendarDates)

	expected := []*CalendarDate{
		{
			ServiceID:     "18033099921-20180524131340_v66.89",
			Date:          NewTimestamp("2018-05-17T00:00:00.000Z"),
			ExceptionType: 1,
		},
		{
			ServiceID:     "18033099921-20180524131340_v66.89",
			Date:          NewTimestamp("2018-05-20T00:00:00.000Z"),
			ExceptionType: 1,
		},
	}

	assert.Equal(expected, calendarDates)
}
