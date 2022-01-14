package gormx

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type TimeStamp struct {
	time.Time
}

func (u TimeStamp) MarshalJSON() ([]byte, error) {
	ts := "0"
	if !u.Time.IsZero() {
		ts = fmt.Sprintf("%d", u.Time.Unix())
	}
	return []byte(ts), nil
}

func (u *TimeStamp) Scan(src interface{}) error {
	value, ok := src.(time.Time)
	if ok {
		*u = TimeStamp{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", src)
}

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
