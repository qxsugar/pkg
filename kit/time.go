package kit

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

// TimeStamp is a custom time type that marshals to Unix timestamp in JSON
// and implements database scanner and valuer interfaces.
type TimeStamp struct {
	time.Time
}

// MarshalJSON implements json.Marshaler, converting time to Unix timestamp.
func (u TimeStamp) MarshalJSON() ([]byte, error) {
	ts := "0"
	if !u.Time.IsZero() {
		ts = fmt.Sprintf("%d", u.Time.Unix())
	}
	return []byte(ts), nil
}

// Scan implements sql.Scanner interface for reading time from database.
func (u *TimeStamp) Scan(src any) error {
	value, ok := src.(time.Time)
	if ok {
		*u = TimeStamp{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", src)
}

// Value implements driver.Valuer interface for writing time to database.
func (u TimeStamp) Value() (driver.Value, error) {
	var zeroTime time.Time
	if u.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return u.Time, nil
}

var _ driver.Valuer = (*TimeStamp)(nil)
var _ sql.Scanner = (*TimeStamp)(nil)
var _ json.Marshaler = (*TimeStamp)(nil)
