package test

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	at "github.com/Mungrel/at-go"
)

type mockClient struct{}

func newMockClient() *at.Client {
	return &at.Client{
		APIKey: "dummy-key",
		Client: &mockClient{},
	}
}

func (mc *mockClient) Do(req *http.Request) (*http.Response, error) {
	url := req.URL.Path

	if strings.Contains(url, "/v2/gtfs/agency") {
		return newResponse(http.StatusOK, getJSON(agencies)), nil
	}

	if strings.Contains(url, "/v2/gtfs/calendar") {
		return newResponse(http.StatusOK, getJSON(calendars)), nil
	}

	if strings.Contains(url, "/v2/gtfs/calendar/serviceId/") {
		return newResponse(http.StatusOK, getJSON(calendarByService)), nil
	}

	return nil, errors.New("Endpoint not mocked")
}

func newResponse(code int, body string) *http.Response {
	return &http.Response{
		StatusCode: code,
		Body:       ioutil.NopCloser(strings.NewReader(body)),
	}
}
