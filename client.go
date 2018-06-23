package at

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const baseURL = "https://api.at.govt.nz/v2"

type clientInterface interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client is a basic AT API client
type Client struct {
	APIKey string
	Client clientInterface
}

// NewATClient creates a new Client
func NewATClient(apiKey string) *Client {
	return &Client{
		APIKey: apiKey,
		Client: http.DefaultClient,
	}
}

func (c *Client) get(url string, result interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(respBytes, &result)
}
