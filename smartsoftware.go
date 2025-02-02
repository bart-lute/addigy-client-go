package addigy

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (c *Client) GetAllSmartSoftware() error {
	rb, err := json.Marshal(c.PagerInfo)
	if err != nil {
		return err
	}
	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/oa/smart-software/query", c.HostUrl), strings.NewReader(string(rb)))
	if err != nil {
		return err
	}

	body, err := c.doRequest(request)
	if err != nil {
		return err
	}
	fmt.Println(string(body))
	return nil
}
