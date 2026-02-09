package addigy_client_go

import (
	"net/http"

	"github.com/bart-lute/addigy-client-go/models/device_entities"
)

const perPage = 50

func (c *Client) Devices(desiredFactIdentifiers *[]string, queryFilter *device_entities.QueryFilter, sortDirection string, sortField string) *[]device_entities.DeviceAudit {

	var items []device_entities.DeviceAudit

	deviceFilter := device_entities.DeviceFilter{
		DesiredFactIdentifiers: *desiredFactIdentifiers,
		Page:                   1,
		PerPage:                perPage,
		Query:                  *queryFilter,
		SortDirection:          sortDirection,
		SortField:              sortField,
	}

	for {
		var deviceAuditResponse device_entities.DeviceAuditResponse
		c.doRequest(http.MethodPost, "devices/", &deviceFilter, &deviceAuditResponse)
		items = append(items, deviceAuditResponse.Items...)
		if deviceAuditResponse.Metadata.Page == deviceAuditResponse.Metadata.PageCount {
			break
		}
		deviceFilter.Page++
	}

	return &items
}
