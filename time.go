package at

import (
	"fmt"
	"strconv"
	"time"
)

const layout = "2018-05-17T00:00:00.000Z"

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
	ts := time.Time(*t).Unix()
	stamp := fmt.Sprint(ts)

	return []byte(stamp), nil
}

// UnmarshalJSON unmarshals JSON into a Timestamp
func (t *Timestamp) UnmarshalJSON(b []byte) error {
	ts, err := strconv.Atoi(string(b))
	if err != nil {
		return err
	}

	*t = Timestamp(time.Unix(int64(ts), 0))

	return nil
}
