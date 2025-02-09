package ade_service

type AdeAccount struct {
	AdminId    string   `json:"admin_id"`
	OrgAddress string   `json:"org_address"`
	OrgEmail   string   `json:"org_email"`
	OrgName    string   `json:"org_name"`
	OrgPhone   string   `json:"org_phone"`
	ServerName string   `json:"server_name"`
	ServerUuid string   `json:"server_uuid"`
	Urls       []string `json:"urls"`
}
