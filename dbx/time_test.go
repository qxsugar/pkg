package dbx

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestMarshalJSON(t *testing.T) {
	ts := TimeStamp{Time: time.Unix(1631295609, 0)}
	b, err := json.Marshal(ts)
	assert.NoError(t, err)
	assert.Equal(t, []byte("1631295609"), b)
}

func TestScan(t *testing.T) {
	var ts TimeStamp
	err := ts.Scan(time.Now())
	assert.NoError(t, err)

	// Test invalid Scan
	err = ts.Scan("invalid")
	assert.Error(t, err)
}

func TestValue(t *testing.T) {
	ts := TimeStamp{Time: time.Unix(1631295609, 0)}
	val, err := ts.Value()
	assert.NoError(t, err)
	assert.IsType(t, ts.Time, val)

	// Test zero time
	zeroTs := TimeStamp{Time: time.Time{}}
	val, err = zeroTs.Value()
	assert.NoError(t, err)
	assert.Nil(t, val)
}
