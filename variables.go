package addigy

import "encoding/json"

func (c *Client) getVariablesItems(url string, queryRequest *QueryRequest, variablesItems *[]VariablesItem) error {
	jsonItems, err := c.getJsonItems(url, queryRequest)
	if err != nil {
		return err
	}

	err = json.Unmarshal(*jsonItems, variablesItems)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) GetAllVariables() (*[]VariablesItem, error) {
	return c.GetVariables(&VariablesQuery{})
}

func (c *Client) GetVariables(query *VariablesQuery) (*[]VariablesItem, error) {

	variablesQueryRequest := QueryRequest{
		PerPage:       perPage,
		SortDirection: sortDirection,
		SortField:     sortField,
		Query:         query,
	}
	url := "oa/variables/query"

	var variablesItems []VariablesItem
	err := c.getVariablesItems(url, &variablesQueryRequest, &variablesItems)
	if err != nil {
		return nil, err
	}

	return &variablesItems, nil
}
