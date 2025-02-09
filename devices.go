package addigy

import (
	"encoding/json"
	"github.com/bart-lute/addigy-client-go/models/device_enties"
)

func (c *Client) getDevicesItems(url string, queryRequest *QueryRequest, deviceAudits *[]device_enties.DeviceAudit) error {
	jsonItems, err := c.getJsonItems(url, queryRequest)
	if err != nil {
		return err
	}

	err = json.Unmarshal(*jsonItems, deviceAudits)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) GetAllDevices(desiredFactIdentifiers *[]string) (*[]device_enties.DeviceAudit, error) {
	return c.GetDevices(&device_enties.QueryFilter{}, desiredFactIdentifiers)
}

func (c *Client) GetDevices(queryFilter *device_enties.QueryFilter, desiredFactIdentifiers *[]string) (*[]device_enties.DeviceAudit, error) {

	devicesQueryRequest := QueryRequest{
		DesiredFactIdentifiers: *desiredFactIdentifiers,
		PerPage:                perPage,
		SortDirection:          sortDirection,
		SortField:              "device_name",
		Query:                  queryFilter,
	}
	url := "devices"

	var deviceAudits []device_enties.DeviceAudit
	err := c.getDevicesItems(url, &devicesQueryRequest, &deviceAudits)
	if err != nil {
		return nil, err
	}

	return &deviceAudits, nil
}
