package kit

import (
	"database/sql/driver"
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

		// Test with nil JSON
		var nilJSONPtr *JSON
		bytes, err = json.Marshal(nilJSONPtr)
		assert.NoError(t, err)
		assert.Equal(t, "null", string(bytes))

		// Test with empty JSON
		emptyJSON := JSON("")
		bytes, err = json.Marshal(emptyJSON)
		assert.NoError(t, err)
		assert.Equal(t, "null", string(bytes))
	})

	t.Run("Scan", func(t *testing.T) {
		j := &JSON{}
		err := j.Scan([]byte(`{"key":"value"}`))
		assert.NoError(t, err)
		assert.Equal(t, `{"key":"value"}`, string(*j))

		// Test scanning invalid string
		err = j.Scan("invalid")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to unmarshal JSONB value")

		// Test scanning nil
		err = j.Scan(nil)
		assert.NoError(t, err)
		assert.Equal(t, JSON{}, *j)

		// Test scanning empty byte slice
		err = j.Scan([]byte{})
		assert.NoError(t, err)
		assert.Equal(t, JSON{}, *j)

		// Test scanning invalid JSON
		err = j.Scan([]byte(`{"invalid": json}`))
		assert.Error(t, err)

		// Test scanning valid JSON array
		err = j.Scan([]byte(`[1,2,3]`))
		assert.NoError(t, err)
		assert.Equal(t, `[1,2,3]`, string(*j))

		// Test scanning JSON null
		err = j.Scan([]byte(`null`))
		assert.NoError(t, err)
		assert.Equal(t, `null`, string(*j))
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

		// Test with complex JSON
		complexJSON := JSON(`{"users":[{"id":1,"name":"John"},{"id":2,"name":"Jane"}],"count":2}`)
		value, err = complexJSON.Value()
		assert.NoError(t, err)
		expected := `{"users":[{"id":1,"name":"John"},{"id":2,"name":"Jane"}],"count":2}`
		assert.Equal(t, expected, string(value.([]byte)))

		// Test with array JSON
		arrayJSON := JSON(`[1,2,3,4,5]`)
		value, err = arrayJSON.Value()
		assert.NoError(t, err)
		assert.Equal(t, `[1,2,3,4,5]`, string(value.([]byte)))
	})
}

func TestJSON_DatabaseInterfaces(t *testing.T) {
	t.Run("implements driver.Valuer", func(t *testing.T) {
		var j JSON = JSON(`{"test":"value"}`)
		var valuer driver.Valuer = &j

		value, err := valuer.Value()
		assert.NoError(t, err)
		assert.Equal(t, `{"test":"value"}`, string(value.([]byte)))
	})

	t.Run("implements sql.Scanner", func(t *testing.T) {
		j := &JSON{}

		// Test scanning as sql.Scanner
		err := j.Scan([]byte(`{"scanned":"data"}`))
		assert.NoError(t, err)
		assert.Equal(t, `{"scanned":"data"}`, string(*j))
	})

	t.Run("implements json.Marshaler", func(t *testing.T) {
		j := JSON(`{"marshaler":"test"}`)
		var marshaler json.Marshaler = j

		bytes, err := marshaler.MarshalJSON()
		assert.NoError(t, err)
		assert.Equal(t, `{"marshaler":"test"}`, string(bytes))
	})
}

func TestJSON_RoundTrip(t *testing.T) {
	t.Run("scan and value round trip", func(t *testing.T) {
		original := `{"id":123,"name":"test","nested":{"key":"value"},"array":[1,2,3]}`

		// Scan the original data
		j := &JSON{}
		err := j.Scan([]byte(original))
		assert.NoError(t, err)

		// Get the value back
		value, err := j.Value()
		assert.NoError(t, err)

		// Should match the original
		assert.Equal(t, original, string(value.([]byte)))
	})

	t.Run("marshal and unmarshal round trip", func(t *testing.T) {
		original := JSON(`{"test":"round trip"}`)

		// Marshal to JSON - this will wrap the JSON in quotes
		marshaled, err := json.Marshal(original)
		assert.NoError(t, err)

		// The result should be the JSON string itself (without extra quotes)
		assert.Equal(t, `{"test":"round trip"}`, string(marshaled))
	})
}

func TestJSON_EdgeCases(t *testing.T) {
	t.Run("very large JSON", func(t *testing.T) {
		// Create a large JSON object with valid content
		largeValue := make([]byte, 1000)
		for i := range largeValue {
			largeValue[i] = 'x'
		}
		largeData := `{"data":"` + string(largeValue) + `"}`

		j := &JSON{}
		err := j.Scan([]byte(largeData))
		assert.NoError(t, err)

		value, err := j.Value()
		assert.NoError(t, err)
		assert.Equal(t, largeData, string(value.([]byte)))
	})

	t.Run("special characters in JSON", func(t *testing.T) {
		specialChars := `{"unicode":"测试","emoji":"😀","quote":"\"quoted\"","newline":"line1\nline2"}`

		j := &JSON{}
		err := j.Scan([]byte(specialChars))
		assert.NoError(t, err)

		value, err := j.Value()
		assert.NoError(t, err)
		assert.Equal(t, specialChars, string(value.([]byte)))
	})

	t.Run("nested objects", func(t *testing.T) {
		nested := `{"level1":{"level2":{"level3":{"value":"deep"}}}}`

		j := &JSON{}
		err := j.Scan([]byte(nested))
		assert.NoError(t, err)

		value, err := j.Value()
		assert.NoError(t, err)
		assert.Equal(t, nested, string(value.([]byte)))
	})
}
