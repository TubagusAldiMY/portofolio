package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
)

// StringList stores a list of strings in MySQL JSON column.
// It also tolerates legacy CSV strings when reading existing data.
type StringList []string

func (s *StringList) Scan(value any) error {
	if value == nil {
		*s = nil
		return nil
	}
	switch v := value.(type) {
	case []byte:
		// Try JSON array first
		var arr []string
		if err := json.Unmarshal(v, &arr); err == nil {
			*s = arr
			return nil
		}
		// Try JSON string
		var str string
		if err := json.Unmarshal(v, &str); err == nil {
			*s = splitCSV(str)
			return nil
		}
		// Fallback: treat as raw string
		*s = splitCSV(string(v))
		return nil
	case string:
		*s = splitCSV(v)
		return nil
	default:
		return fmt.Errorf("unsupported Scan type %T", value)
	}
}

func (s StringList) Value() (driver.Value, error) {
	// Store as JSON array string
	b, err := json.Marshal([]string(s))
	if err != nil {
		return nil, err
	}
	return string(b), nil
}

func splitCSV(in string) []string {
	parts := strings.Split(in, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			out = append(out, p)
		}
	}
	return out
}
