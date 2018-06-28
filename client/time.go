package client

import (
	"time"
)

const layout = "2006-01-02T15:04:05.000Z"

// Timestamp is an alias for time.Time
type Timestamp time.Time

// NewTimestamp is a convenience method for creating a *Timestamp
func NewTimestamp(dateTime string) *Timestamp {
	t, err := time.Parse(layout, dateTime)
	if err != nil {
		panic(err)
	}

	stamp := Timestamp(t)
	return &stamp
}

// MarshalJSON marshals a Timestamp into JSON
func (t *Timestamp) MarshalJSON() ([]byte, error) {
	ts := time.Time(*t).Format(layout)
	return []byte(ts), nil
}

// UnmarshalJSON unmarshals JSON into a Timestamp
func (t *Timestamp) UnmarshalJSON(b []byte) error {
	b = b[1 : len(b)-1]
	ts, err := time.Parse(layout, string(b))
	if err != nil {
		return err
	}

	*t = Timestamp(ts)

	return nil
}
