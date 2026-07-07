package atc

import (
	"encoding/json"
	"fmt"
)

type HTTPHeadersFlag map[string]string

func (h *HTTPHeadersFlag) UnmarshalFlag(value string) error {
	var headers map[string]string
	if err := json.Unmarshal([]byte(value), &headers); err != nil {
		return fmt.Errorf("invalid JSON for additional-http-headers (expected {\"Name\":\"Value\"}): %s", err)
	}
	*h = headers
	return nil
}

