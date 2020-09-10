package matchers

import (
	"fmt"
	"regexp"
)

type regexpMatcher struct {
	pattern *regexp.Regexp
}

func Regexp(pattern string) *regexpMatcher {
	return &regexpMatcher{
		pattern: regexp.MustCompile(pattern),
	}
}

func (m *regexpMatcher) String() string {
	return fmt.Sprintf("matches pattern /%v/", m.pattern)
}

func (m *regexpMatcher) Matches(x interface{}) bool {
	s, ok := x.(string)
	if !ok {
		return false
	}

	return m.pattern.MatchString(s)
}
