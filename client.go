package addigy

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

const HostUrl string = "https://api.addigy.com/api/v2"

type Client struct {
	HostUrl    string
	HTTPClient *http.Client
	ApiKey     string
	PagerInfo  struct {
		Page          int    `json:"page" default:"1"`
		PerPage       int    `json:"per_page" default:"50"`
		SortDirection string `json:"sort_direction" default:"asc"`
		SortField     string `json:"sort_field" default:"name"`
	}
}

func NewClient(host, apiKey *string) *Client {
	c := Client{
		HostUrl:    HostUrl,
		ApiKey:     *apiKey,
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
	}

	if host != nil {
		c.HostUrl = *host
	}

	return &c
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", c.ApiKey)

	response, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code: %d, body: %s", response.StatusCode, string(body))
	}

	return body, nil
}
