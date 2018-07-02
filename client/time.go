package client

import (
	"time"
)

var layouts = []string{
	"2006-01-02T15:04:05.000Z",
	"2006-01-02T15:04:05Z",
	"2006-01-02T15:04:05",
}

// Timestamp is an alias for time.Time
type Timestamp time.Time

// NewTimestamp is a convenience method for creating a *Timestamp
func NewTimestamp(dateTime string) *Timestamp {
	t, err := parseTime(dateTime)
	if err != nil {
		panic(err)
	}

	stamp := Timestamp(t)
	return &stamp
}

// UnmarshalJSON unmarshals JSON into a Timestamp
func (t *Timestamp) UnmarshalJSON(b []byte) error {
	b = b[1 : len(b)-1]
	ts, err := parseTime(string(b))
	if err != nil {
		return err
	}

	*t = Timestamp(ts)

	return nil
}

func parseTime(timeStr string) (time.Time, error) {
	var err error
	for _, format := range layouts {
		t, err := time.Parse(format, timeStr)
		if err == nil {
			return t, nil
		}
	}
	return time.Time{}, err
}
