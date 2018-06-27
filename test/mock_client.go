package test

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	at "github.com/Mungrel/at-go"
)

const mockServiceID = "123456ABC"

type mockClient struct{}

func newMockClient() *at.Client {
	return &at.Client{
		APIKey: "dummy-key",
		Client: &mockClient{},
	}
}

func (mc *mockClient) Do(req *http.Request) (*http.Response, error) {
	url := req.URL.Path

	if strings.HasSuffix(url, "/v2/gtfs/agency") {
		return newResponse(http.StatusOK, getJSON(agencies)), nil
	}

	if strings.HasSuffix(url, "/v2/gtfs/calendar") {
		return newResponse(http.StatusOK, getJSON(calendars)), nil
	}

	if strings.HasSuffix(url, "/v2/gtfs/calendar/serviceId/"+mockServiceID) {
		return newResponse(http.StatusOK, getJSON(calendarByService)), nil
	}

	if strings.HasSuffix(url, "/v2/gtfs/calendarDate") {
		return newResponse(http.StatusOK, getJSON(calendarDates)), nil
	}

	if strings.HasSuffix(url, "/v2/gtfs/calendarDate/serviceId/"+mockServiceID) {
		return newResponse(http.StatusOK, getJSON(calendarDatesByService)), nil
	}

	if strings.HasSuffix(url, "/v2/gtfs/metadata") {
		return newResponse(http.StatusOK, getJSON(metadata)), nil
	}

	if strings.HasSuffix(url, "/v2/gtfs/stops") {
		return newResponse(http.StatusOK, getJSON(stops)), nil
	}

	if strings.HasSuffix(url, "/v2/gtfs/trips") {
		return newResponse(http.StatusOK, getJSON(trips)), nil
	}

	if strings.HasSuffix(url, "/v2/gtfs/versions") {
		return newResponse(http.StatusOK, getJSON(versions)), nil
	}

	return nil, errors.New("Endpoint not mocked")
}

func newResponse(code int, body string) *http.Response {
	return &http.Response{
		StatusCode: code,
		Body:       ioutil.NopCloser(strings.NewReader(body)),
	}
}
