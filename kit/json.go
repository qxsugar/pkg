package kit

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

// JSON is a custom type for handling JSON data in database operations.
// It implements driver.Valuer, sql.Scanner, and json.Marshaler interfaces
// to provide seamless JSON handling between Go structs and database fields.
type JSON json.RawMessage

// MarshalJSON implements json.Marshaler interface.
func (j JSON) MarshalJSON() ([]byte, error) {
	// null or empty string should be marshaled to null
	if len(j) == 0 {
		return []byte("null"), nil
	}
	return j, nil
}

// Scan implements sql.Scanner interface for reading JSON data from database.
func (j *JSON) Scan(value any) error {
	if value == nil {
		*j = JSON{}
		return nil
	}

	var bytes []byte
	switch v := value.(type) {
	case []byte:
		bytes = v
	case string:
		bytes = []byte(v)
	case nil:
		*j = JSON{}
		return nil
	default:
		return fmt.Errorf("failed to unmarshal JSONB value: unsupported type %T", value)
	}

	if len(bytes) == 0 {
		*j = JSON{}
		return nil
	}

	result := json.RawMessage{}
	err := json.Unmarshal(bytes, &result)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSONB value: %w", err)
	}
	*j = JSON(result)
	return nil
}

// Value implements driver.Valuer interface for writing JSON data to database.
func (j JSON) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}
	return j.MarshalJSON()
}

var _ driver.Valuer = (*JSON)(nil)
var _ sql.Scanner = (*JSON)(nil)
var _ json.Marshaler = (*JSON)(nil)
