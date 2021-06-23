package matchers

import (
	"fmt"
	"regexp"
)

type RegexpMatcher struct {
	pattern *regexp.Regexp
}

func Regexp(pattern string) *RegexpMatcher {
	return &RegexpMatcher{
		pattern: regexp.MustCompile(pattern),
	}
}

func (m *RegexpMatcher) String() string {
	return fmt.Sprintf("matches pattern /%v/", m.pattern)
}

func (m *RegexpMatcher) Matches(x interface{}) bool {
	s, ok := x.(string)
	if !ok {
		return false
	}

	return m.pattern.MatchString(s)
}
