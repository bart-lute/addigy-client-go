package addigy

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"slices"
	"strings"
	"time"
)

const HostUrl string = "https://api.addigy.com/api/v2"

// defaults
var (
	perPage       int    = 100
	sortField     string = "name"
	sortDirection string = "asc"
)

type Client struct {
	HostUrl    string
	HTTPClient *http.Client
	ApiKey     string
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

// Return a Reader object which can be used for the Request
func newQueryRequest(QueryRequest *QueryRequest) (*strings.Reader, error) {
	rb, err := json.Marshal(QueryRequest)
	if err != nil {
		return nil, err
	}
	// bump the page counter to prepare for the next page
	QueryRequest.Page++
	return strings.NewReader(string(rb)), nil
}

func (c *Client) getJsonItems(url string, queryRequest *QueryRequest) (*[]byte, error) {
	var items []any
	queryRequest.Page = 1
	for {
		requestBody, err := newQueryRequest(queryRequest)
		if err != nil {
			return nil, err
		}

		request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/%s", c.HostUrl, url), requestBody)
		if err != nil {
			return nil, err
		}

		body, err := c.doRequest(request)
		if err != nil {
			return nil, err
		}

		var queryResponse QueryResponse
		err = json.Unmarshal(body, &queryResponse)
		if err != nil {
			return nil, err
		}

		items = slices.Concat(items, queryResponse.Items)

		if queryRequest.Page > queryResponse.Metadata.PageCount {
			break
		}
	}

	// convert the generic items slice back to Json and return the pointer
	bytes, err := json.Marshal(items)
	if err != nil {
		return nil, err
	}

	return &bytes, nil
}
