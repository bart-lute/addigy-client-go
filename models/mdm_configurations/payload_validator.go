package mdm_configurations

import "encoding/json"

type PayloadValidator struct {
	Constrains       []Constrain `json:"constrains"`
	ExactString      string      `json:"exact_string"`
	MaxFloat         json.Number `json:"max_float"`
	MinFloat         json.Number `json:"min_float"`
	Regex            string      `json:"regex"`
	Required         bool        `json:"required"`
	SupportedInts    []int       `json:"supported_ints"`
	SupportedStrings []string    `json:"supported_strings"`
	Unique           bool        `json:"unique"`
}
