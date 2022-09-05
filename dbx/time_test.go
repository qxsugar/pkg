package dbx

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func MarshalJSON(t *testing.T) {
	ts := TimeStamp{Time: time.Now()}
	tsValue, err := json.Marshal(ts)
	if err != nil {
		t.Fail()
	}

	if string(tsValue) != fmt.Sprint(ts.Unix()) {
		t.Fail()
	}
}
