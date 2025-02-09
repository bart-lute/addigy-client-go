package addigy

import (
	"encoding/json"
)

func (c *Client) getCustomFactsItems(url string, queryRequest *QueryRequest, customFactsItems *[]CustomFactsItem) error {
	jsonItems, err := c.getJsonItems(url, queryRequest)
	if err != nil {
		return err
	}

	err = json.Unmarshal(*jsonItems, customFactsItems)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) GetAllCustomFacts() (*[]CustomFactsItem, error) {
	return c.GetCustomFacts(&CustomFactsQuery{})
}

func (c *Client) GetCustomFacts(query *CustomFactsQuery) (*[]CustomFactsItem, error) {

	customFactsQueryRequest := QueryRequest{
		PerPage:       perPage,
		SortDirection: sortDirection,
		SortField:     sortField,
		Query:         query,
	}
	url := "facts/custom/query"

	var customSoftwareItems []CustomFactsItem
	err := c.getCustomFactsItems(url, &customFactsQueryRequest, &customSoftwareItems)
	if err != nil {
		return nil, err
	}

	return &customSoftwareItems, nil
}
