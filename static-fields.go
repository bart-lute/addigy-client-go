package addigy_client_go

import (
	"fmt"
	"net/http"

	"github.com/bart-lute/addigy-client-go/models/response_entities"
)

func (c *Client) StaticFieldsValues() *[]response_entities.StaticFieldsValueItem {
	var items []response_entities.StaticFieldsValueItem
	page := 1
	perPage := 50

	for {
		var response response_entities.PaginatedResponseStaticFieldsEntitiesStaticFieldValueResponse
		endPoint := fmt.Sprintf("static-fields/value?page=%d&per_page=%d", page, perPage)
		c.doRequest(http.MethodGet, endPoint, nil, &response)
		items = append(items, response.Items...)
		if page >= response.Metadata.PageCount {
			break
		}
		page++
	}

	return &items
}
