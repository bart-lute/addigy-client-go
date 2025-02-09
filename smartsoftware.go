package addigy

import (
	"encoding/json"
)

func (c *Client) getSmartSoftwareItems(url string, queryRequest *QueryRequest, smartSoftwareItems *[]SmartSoftwareItem) error {
	jsonItems, err := c.getJsonItems(url, queryRequest)
	if err != nil {
		return err
	}

	err = json.Unmarshal(*jsonItems, smartSoftwareItems)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) GetAllSmartSoftware() (*[]SmartSoftwareItem, error) {
	return c.GetSmartSoftware(&SmartSoftwareQuery{})
}

func (c *Client) GetSmartSoftware(query *SmartSoftwareQuery) (*[]SmartSoftwareItem, error) {

	smartSoftwareQueryRequest := QueryRequest{
		PerPage:       perPage,
		SortDirection: sortDirection,
		SortField:     sortField,
		Query:         *query,
	}

	var smartSoftwareItems []SmartSoftwareItem
	err := c.getSmartSoftwareItems("oa/smart-software/query", &smartSoftwareQueryRequest, &smartSoftwareItems)
	if err != nil {
		return nil, err
	}

	return &smartSoftwareItems, nil
}
