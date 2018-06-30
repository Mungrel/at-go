package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	baseURL      = "https://api.at.govt.nz"
	apiKeyHeader = "Ocp-Apim-Subscription-Key"
	okStatus     = "OK"
)

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

type response struct {
	Status   string      `json:"status"`
	Response interface{} `json:"response"`
	Error    *string     `json:"error"`
}

func (c *Client) get(url string, result interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.Header.Set(apiKeyHeader, c.APIKey)

	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	respStruct := response{
		Response: result,
	}

	err = json.Unmarshal(respBytes, &respStruct)
	if err != nil {
		return err
	}

	if respStruct.Error != nil {
		return errors.New(*respStruct.Error)
	}

	if respStruct.Status != okStatus {
		return fmt.Errorf("Non OK status response: %s", respStruct.Status)
	}

	return nil
}
