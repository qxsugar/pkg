package kit

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJSON(t *testing.T) {
	t.Run("MarshalJSON", func(t *testing.T) {
		j := JSON(`{"key":"value"}`)
		bytes, err := json.Marshal(j)
		assert.NoError(t, err)
		assert.Equal(t, `{"key":"value"}`, string(bytes))

		var nilJSON = JSON{}
		bytes, err = json.Marshal(nilJSON)
		assert.NoError(t, err)
		assert.Equal(t, "null", string(bytes))
	})

	t.Run("Scan", func(t *testing.T) {
		j := &JSON{}
		err := j.Scan([]byte(`{"key":"value"}`))
		assert.NoError(t, err)
		assert.Equal(t, `{"key":"value"}`, string(*j))

		err = j.Scan("invalid")
		assert.Error(t, err)
	})

	t.Run("Value", func(t *testing.T) {
		j := JSON(`{"key":"value"}`)
		value, err := j.Value()
		assert.NoError(t, err)
		assert.Equal(t, `{"key":"value"}`, string(value.([]byte)))

		var emptyJSON JSON
		value, err = emptyJSON.Value()
		assert.NoError(t, err)
		assert.Nil(t, value)
	})
}
