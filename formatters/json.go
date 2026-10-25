package formatters

import (
	"encoding/json"

	"code/internal/diff"
)

func formatJSON(tree []diff.Node) (string, error) {
	result, err := json.MarshalIndent(tree, "", "  ")
	if err != nil {
		return "", err
	}
	return string(result), nil
}
