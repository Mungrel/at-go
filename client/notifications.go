package client

import (
	"encoding/json"
	"errors"
	"strconv"
)

// notifications represents a notifications response from the AT Notifications API
type notifications struct {
	Data []*Notification `json:"data"`
	Time *Timestamp      `json:"time"`
}

// Notification represents a notification in the AT Notifications API
type Notification struct {
	ID                 string     `json:"id"`
	OrderingDate       *Timestamp `json:"orderingDate"`
	Type               string     `json:"type"`
	Title              string     `json:"title"`
	StopID             int64      `json:"stopId"`
	ExistingAddress    string     `json:"existingAddress"`
	Description        string     `json:"description"`
	Impact             string     `json:"impact"`
	StartDate          *Timestamp `json:"startDate"`
	EndDate            *Timestamp `json:"endDate"`
	ExpectedResolution string     `json:"expectedResolution"`
	Location           *Location  `json:"location"`
}

// Location represnets a notification Location
type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Address   string  `json:"address"`
}

type notificationInterfaceID struct {
	ID                 interface{} `json:"id"`
	OrderingDate       *Timestamp  `json:"orderingDate"`
	Type               string      `json:"type"`
	Title              string      `json:"title"`
	StopID             int64       `json:"stopId"`
	ExistingAddress    string      `json:"existingAddress"`
	Description        string      `json:"description"`
	Impact             string      `json:"impact"`
	StartDate          *Timestamp  `json:"startDate"`
	EndDate            *Timestamp  `json:"endDate"`
	ExpectedResolution string      `json:"expectedResolution"`
	Location           *Location   `json:"location"`
}

func (n *notificationInterfaceID) toStrID() *Notification {
	var idStr string
	idStr, ok := n.ID.(string)
	if !ok {
		id, ok := n.ID.(float64)
		if !ok {
			panic(errors.New("Could not cast ID to float64 or string"))
		}
		idStr = strconv.FormatInt(int64(id), 10)
		// idStr = strconv.FormatInt(id, 64)
	}
	return &Notification{
		ID:                 idStr,
		OrderingDate:       n.OrderingDate,
		Type:               n.Type,
		Title:              n.Title,
		StopID:             n.StopID,
		ExistingAddress:    n.ExistingAddress,
		Description:        n.Description,
		Impact:             n.Impact,
		StartDate:          n.StartDate,
		EndDate:            n.EndDate,
		ExpectedResolution: n.ExpectedResolution,
		Location:           n.Location,
	}
}

// UnmarshalJSON unmarshals notifications, as the ID field may be int64 or string
func (n *Notification) UnmarshalJSON(b []byte) error {
	var notification notificationInterfaceID

	err := json.Unmarshal(b, &notification)
	if err != nil {
		return err
	}

	*n = *notification.toStrID()
	return nil
}

// Notifications gets a list of Notifications
func (client *Client) Notifications() ([]*Notification, error) {
	url := baseURL + "/v2/notifications/"

	var response notifications
	err := client.get(url, &response)

	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

// NotificationsByCategory gets a list of Notifications, filtered by category
func (client *Client) NotificationsByCategory(category string) ([]*Notification, error) {
	url := baseURL + "/v2/notifications/" + category

	var response notifications
	err := client.get(url, &response)

	if err != nil {
		return nil, err
	}

	return response.Data, nil
}
