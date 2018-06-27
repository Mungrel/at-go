package test

import "io/ioutil"

const (
	partialPath            = "payloads/"
	agencies               = "agencies.json"
	calendars              = "calendars.json"
	calendarByService      = "calendar_by_service.json"
	calendarDates          = "calendar_dates.json"
	calendarDatesByService = "calendar_dates_by_service.json"
	metadata               = "metadata.json"
	routes                 = "routes.json"
	stops                  = "stops.json"
	trips                  = "trips.json"
	versions               = "versions.json"
)

func getJSON(fileName string) string {
	bytes, err := ioutil.ReadFile(partialPath + fileName)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}
