package mdm_commands

type FileVaultUnlock struct {
	Password                 string      `json:"Password"`
	PrivateKeyExport         interface{} `json:"PrivateKeyExport"` // Check Later
	PrivateKeyExportPassword string      `json:"PrivateKeyExportPassword"`
}
type RotateFileVaultKeyCommandExample struct {
	FileVaultUnlock              FileVaultUnlock `json:"FileVaultUnlock"`
	KeyType                      string          `json:"KeyType"`
	NewCertificate               interface{}     `json:"NewCertificate"`
	ReplyEncryptionCertificate   string          `json:"ReplyEncryptionCertificate"`
	RequestRequiresNetworkTether bool            `json:"RequestRequiresNetworkTether"`
	RequestType                  string          `json:"RequestType"`
}
