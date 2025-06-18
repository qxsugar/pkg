package kit

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestWrapLike(t *testing.T) {
	result := WrapLike("test")
	expected := "%test%"
	assert.Equal(t, expected, result)
}

func TestWrapLeftLike(t *testing.T) {
	result := WrapLeftLike("test")
	expected := "%test"
	assert.Equal(t, expected, result)
}
