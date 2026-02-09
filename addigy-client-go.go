package addigy_client_go

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"strings"
	"time"
)

type Client struct {
	baseUrl    string
	apiKey     string
	httpClient *http.Client
}

// Timeout for API calls in seconds
const timeout = 10 * time.Second

func Init(baseUrl string, apiKey string) *Client {
	httpClient := &http.Client{Timeout: 10 * time.Second}
	return &Client{
		baseUrl:    baseUrl,
		apiKey:     apiKey,
		httpClient: httpClient,
	}
}

func (c *Client) doRequest(method string, endPoint string, requestBody any, responseBody any) {

	slog.Debug(fmt.Sprintf("requesting: %s, method: %s", endPoint, method))

	// Convert the Request Body to a Reader Object
	var requestBodyReader io.Reader
	if requestBody != nil {
		rb, err := json.Marshal(requestBody)
		if err != nil {
			log.Fatal(err)
		}
		requestBodyReader = strings.NewReader(string(rb))
	}

	url := fmt.Sprintf("%s/%s", c.baseUrl, endPoint)
	request, err := http.NewRequest(method, url, requestBodyReader)
	if err != nil {
		log.Fatal(err)
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("x-api-key", c.apiKey)

	response, err := c.httpClient.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(response.Body)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	if response.StatusCode != http.StatusOK {
		log.Fatal(fmt.Sprintf("API error: %s", response.Status))
	}

	if err := json.Unmarshal(body, responseBody); err != nil {
		log.Fatal(err)
	}
}
