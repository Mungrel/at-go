package test

import "io/ioutil"

const (
	partialPath       = "payloads/"
	agencies          = "agencies.json"
	calendars         = "calendars.json"
	calendarByService = "calendar_by_service.json"
)

func getJSON(fileName string) string {
	bytes, err := ioutil.ReadFile(partialPath + fileName)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}
