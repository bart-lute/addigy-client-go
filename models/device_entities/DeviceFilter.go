package device_entities

type DeviceFilter struct {
    DesiredFactIdentifiers []string    `json:"desired_fact_identifiers"`
    Page                   int         `json:"page"`
    PerPage                int         `json:"per_page"`
    Query                  QueryFilter `json:"query"`
    SortDirection          string      `json:"sort_direction"` // asc | desc
    SortField              string      `json:"sort_field"`
}
