package flag

import (
	"fmt"
	"os"

	"go.yaml.in/yaml/v3"
	"golang.org/x/net/http/httpguts"
)

type CustomHTTPHeaders struct {
	Path    string
	Headers map[string]string
}

func (f *CustomHTTPHeaders) UnmarshalFlag(value string) error {
	content, err := os.ReadFile(value)
	if err != nil {
		return fmt.Errorf("Failed to open custom HTTP headers file (%s): %w", value, err)
	}

	var headers map[string]string
	if err = yaml.Unmarshal(content, &headers); err != nil {
		return fmt.Errorf("Failed to parse custom HTTP headers file (%s): %w", value, err)
	}

	if headers == nil {
		headers = map[string]string{}
	}

	for name, val := range headers {
		if !httpguts.ValidHeaderFieldName(name) {
			return fmt.Errorf("Invalid header name %q in custom HTTP headers file (%s)", name, value)
		}
		if !httpguts.ValidHeaderFieldValue(val) {
			return fmt.Errorf("Invalid header value for %q in custom HTTP headers file (%s)", name, value)
		}
	}

	f.Path = value
	f.Headers = headers
	return nil
}