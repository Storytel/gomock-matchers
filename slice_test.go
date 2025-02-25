package matchers_test

import (
	matchers "github.com/Storytel/gomock-matchers/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSliceMatcher(t *testing.T) {
	data := []string{"a", "b", "c"}

	assert := assert.New(t)
	m := matchers.SliceLength[string](3)
	assert.True(m.Matches(data))
}

func TestStructSliceMatcher(t *testing.T) {
	type A struct {
		B int
	}
	data := []A{{1}, {2}, {3}}

	assert := assert.New(t)
	m := matchers.SliceLength[A](3)
	assert.True(m.Matches(data))
}

func TestSliceMatcherNumberOfElementsMismatch(t *testing.T) {
	data := []string{"a", "b", "c"}

	assert := assert.New(t)
	m := matchers.SliceLength[string](2)
	assert.False(m.Matches(data))
}

func TestSliceMatcherString(t *testing.T) {
	assert := assert.New(t)
	m := matchers.SliceLength[int](4)
	assert.Equal("4", m.String())
}
