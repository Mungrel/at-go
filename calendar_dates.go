package at

import "errors"

type calendarDatesResponse struct {
	Status   string          `json:"status"`
	Response []*CalendarDate `json:"response"`
	Error    *string         `json:"error"`
}

// CalendarDate represents an exception to a service as defined in the calendar list in the AT GTFS API
type CalendarDate struct {
	ServiceID     string     `json:"service_id"`
	Date          *Timestamp `json:"date"`
	ExceptionType int        `json:"exception_type"`
}

// CalendarDates gets a list of exceptions to services from the AT GTFS API
func (client *Client) CalendarDates() ([]*CalendarDate, error) {
	url := baseURL + "/v2/gtfs/calendarDate"

	var response calendarDatesResponse
	err := client.get(url, &response)

	if err != nil {
		return nil, err
	}

	if response.Error != nil {
		return nil, errors.New(*response.Error)
	}

	return response.Response, nil
}
