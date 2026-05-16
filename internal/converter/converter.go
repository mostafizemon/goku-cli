package converter

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

// DetectFormat returns "json", "yaml", or "" based on file extension
func DetectFormat(filePath string) string {
	ext := strings.ToLower(strings.TrimPrefix(filepath.Ext(filePath), "."))
	switch ext {
	case "json":
		return "json"
	case "yaml", "yml":
		return "yaml"
	default:
		return ""
	}
}

// Convert takes raw bytes and converts from inputFmt to outputFmt
func Convert(data []byte, inputFmt, outputFmt string) ([]byte, error) {
	// Step 1: Unmarshal input into a generic map
	var intermediate interface{}

	switch inputFmt {
	case "json":
		if err := json.Unmarshal(data, &intermediate); err != nil {
			return nil, fmt.Errorf("invalid JSON: %w", err)
		}
	case "yaml", "yml":
		if err := yaml.Unmarshal(data, &intermediate); err != nil {
			return nil, fmt.Errorf("invalid YAML: %w", err)
		}
	default:
		return nil, fmt.Errorf("unsupported input format: %s", inputFmt)
	}

	// Step 2: Marshal into output format
	switch outputFmt {
	case "json":
		// yaml unmarshals maps as map[string]interface{} but
		// json needs map[string]interface{} — normalize keys
		intermediate = normalizeKeys(intermediate)
		out, err := json.MarshalIndent(intermediate, "", "  ")
		if err != nil {
			return nil, fmt.Errorf("failed to marshal JSON: %w", err)
		}
		return out, nil
	case "yaml", "yml":
		out, err := yaml.Marshal(intermediate)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal YAML: %w", err)
		}
		return out, nil
	default:
		return nil, fmt.Errorf("unsupported output format: %s", outputFmt)
	}
}

// normalizeKeys converts map[interface{}]interface{} (from yaml) to
// map[string]interface{} so JSON marshaling works correctly
func normalizeKeys(v interface{}) interface{} {
	switch val := v.(type) {
	case map[interface{}]interface{}:
		result := make(map[string]interface{})
		for k, v2 := range val {
			result[fmt.Sprintf("%v", k)] = normalizeKeys(v2)
		}
		return result
	case map[string]interface{}:
		for k, v2 := range val {
			val[k] = normalizeKeys(v2)
		}
		return val
	case []interface{}:
		for i, v2 := range val {
			val[i] = normalizeKeys(v2)
		}
		return val
	default:
		return v
	}
}