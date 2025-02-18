package mbov_service

type CatalogCodeUsage struct {
	CatalogCode string `json:"catalog_code"`
	Name        string `json:"name"`
	Usage       int    `json:"usage"`
}
