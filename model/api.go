package model

import "encoding/json"

type OHLCResponse struct {
	Error  []string                   `json:"error"`
	Result map[string]json.RawMessage `json:"result"`
}
