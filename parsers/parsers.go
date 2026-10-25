// Package parsers decodes configuration data without reading or writing files.
package parsers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"gopkg.in/yaml.v3"
)

// Parse decodes one JSON or YAML object into the same set of Go value types.
func Parse(data []byte, format string) (map[string]any, error) {
	switch format {
	case "json":
		return parseJSON(data)
	case "yaml", "yml":
		return parseYAML(data)
	default:
		return nil, fmt.Errorf("unsupported input format %q (expected json, yaml or yml)", format)
	}
}

func parseJSON(data []byte) (map[string]any, error) {
	var result map[string]any
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	if result == nil {
		return nil, errors.New("configuration must be an object")
	}
	return result, nil
}

func parseYAML(data []byte) (map[string]any, error) {
	decoder := yaml.NewDecoder(bytes.NewReader(data))
	var result map[string]any
	if err := decoder.Decode(&result); err != nil {
		return nil, err
	}
	var extra any
	if err := decoder.Decode(&extra); !errors.Is(err, io.EOF) {
		return nil, errors.New("configuration must contain exactly one YAML document")
	}
	// A JSON round trip normalizes YAML numbers and validates nested object keys.
	normalized, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}
	return parseJSON(normalized)
}
