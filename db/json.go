package db

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
)

type JSON json.RawMessage

func (j *JSON) MarshalJSON() ([]byte, error) {
	return *j, nil
}

func (j *JSON) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	result := json.RawMessage{}
	err := json.Unmarshal(bytes, &result)
	*j = JSON(result)
	return err
}

func (j JSON) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}
	return json.RawMessage(j).MarshalJSON()
}

var _ driver.Valuer = (*JSON)(nil)
var _ sql.Scanner = (*JSON)(nil)
var _ json.Marshaler = (*JSON)(nil)
