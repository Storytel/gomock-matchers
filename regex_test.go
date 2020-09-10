package matchers_test

import (
	"fmt"
	"testing"

	matchers "github.com/Storytel/gomock-matchers"
	"github.com/stretchr/testify/assert"
)

func TestRegexpMatcher(t *testing.T) {

	testCases := []struct {
		pattern string
		value   interface{}
		matches bool
	}{
		{"^something .*", "something good", true},
		{"^something .*", "that's something good", false},
		{"^[0-9][1-5]a?$", "42a", true},
		{".*", 12, false},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("Test #%d", i), func(t *testing.T) {
			assert := assert.New(t)
			m := matchers.Regexp(testCase.pattern)
			assert.Equal(testCase.matches, m.Matches(testCase.value))
		})
	}
}

func TestRegexpMatcherCompileError(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("invalid regexp did not panic")
		}
	}()
	_ = matchers.Regexp("^(abc$")
}

func TestRegexpMatcherString(t *testing.T) {
	testCases := []struct {
		pattern string
		str     string
	}{
		{"^[a-z]$", "matches pattern /^[a-z]$/"},
		{"asdf", "matches pattern /asdf/"},
		{"something[[:space:]]", "matches pattern /something[[:space:]]/"},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("Test #%d", i), func(t *testing.T) {
			assert := assert.New(t)
			m := matchers.Regexp(testCase.pattern)
			assert.Equal(testCase.str, m.String())
		})
	}
}
