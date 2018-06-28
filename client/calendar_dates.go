package client

// CalendarDate represents an exception to a service as defined in the calendar list in the AT GTFS API
type CalendarDate struct {
	ServiceID     string     `json:"service_id"`
	Date          *Timestamp `json:"date"`
	ExceptionType int        `json:"exception_type"`
}

// CalendarDates gets a list of exceptions to services from the AT GTFS API
func (client *Client) CalendarDates() ([]*CalendarDate, error) {
	url := baseURL + "/v2/gtfs/calendarDate"

	var response []*CalendarDate
	err := client.get(url, &response)

	if err != nil {
		return nil, err
	}

	return response, nil
}

// CalendarDateByService gets a list of exceptions for a given service ID from the AT GTFS API
func (client *Client) CalendarDateByService(serviceID string) ([]*CalendarDate, error) {
	url := baseURL + "/v2/gtfs/calendarDate/serviceId/" + serviceID

	var response []*CalendarDate
	err := client.get(url, &response)

	if err != nil {
		return nil, err
	}

	return response, nil
}
