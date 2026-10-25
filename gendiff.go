// Package code compares JSON and YAML configuration files.
package code

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"code/formatters"
	"code/internal/diff"
	"code/parsers"
)

// GenDiff returns the difference between two files in stylish, plain or json format.
// An empty format selects stylish. The returned string has no trailing newline.
func GenDiff(filepath1, filepath2, format string) (string, error) {
	first, err := readConfig(filepath1)
	if err != nil {
		return "", err
	}
	second, err := readConfig(filepath2)
	if err != nil {
		return "", err
	}
	return formatters.Format(diff.Build(first, second), format)
}

func readConfig(path string) (map[string]any, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read %q: %w", path, err)
	}
	format := strings.ToLower(strings.TrimPrefix(filepath.Ext(path), "."))
	result, err := parsers.Parse(data, format)
	if err != nil {
		return nil, fmt.Errorf("parse %q: %w", path, err)
	}
	return result, nil
}
