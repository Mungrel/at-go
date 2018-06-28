package client

import "errors"

type calendarsResponse struct {
	Status   string      `json:"status"`
	Response []*Calendar `json:"response"`
	Error    *string     `json:"error"`
}

// Calendar represents a service calendar from the AT GTFS API
type Calendar struct {
	ServiceID string     `json:"service_id"`
	Monday    int        `json:"monday"`
	Tuesday   int        `json:"tuesday"`
	Wednesday int        `json:"wednesday"`
	Thursday  int        `json:"thursday"`
	Friday    int        `json:"friday"`
	Saturday  int        `json:"saturday"`
	Sunday    int        `json:"sunday"`
	StartDate *Timestamp `json:"start_date"`
	EndDate   *Timestamp `json:"end_date"`
}

// Calendars gets a list of service calendars from the AT GTFS API
func (client *Client) Calendars() ([]*Calendar, error) {
	url := baseURL + "/v2/gtfs/calendar"

	var response calendarsResponse
	err := client.get(url, &response)
	if err != nil {
		return nil, err
	}

	if response.Error != nil {
		return nil, errors.New(*response.Error)
	}

	return response.Response, nil
}

// CalendarByService gets a calendar by its serviceID from the AT GTFS API
func (client *Client) CalendarByService(serviceID string) ([]*Calendar, error) {
	url := baseURL + "/v2/gtfs/calendar/serviceId/" + serviceID

	var response calendarsResponse
	err := client.get(url, &response)

	if err != nil {
		return nil, err
	}

	if response.Error != nil {
		return nil, errors.New(*response.Error)
	}

	return response.Response, nil
}
