package kit

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type JSON json.RawMessage

func (j JSON) MarshalJSON() ([]byte, error) {
	// null or empty string should be marshalled to null
	if j == nil || len(j) == 0 {
		return []byte("null"), nil
	}
	return j, nil
}

func (j *JSON) Scan(value any) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to unmarshal JSONB value: %v", value)
	}

	result := json.RawMessage{}
	err := json.Unmarshal(bytes, &result)
	if err != nil {
		return err
	}
	*j = JSON(result)
	return nil
}

func (j *JSON) Value() (driver.Value, error) {
	if len(*j) == 0 {
		return nil, nil
	}
	return j.MarshalJSON()
}

var _ driver.Valuer = (*JSON)(nil)
var _ sql.Scanner = (*JSON)(nil)
var _ json.Marshaler = (*JSON)(nil)
