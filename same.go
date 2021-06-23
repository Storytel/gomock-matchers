package matchers

import "fmt"

type SameMatcher struct {
	x interface{}
}

// Same will return a new Same matcher which uses `==` comparison
func Same(x interface{}) *SameMatcher {
	return &SameMatcher{x}
}

func (m *SameMatcher) String() string {
	return fmt.Sprintf("is same as '%v'", m.x)
}

func (m *SameMatcher) Matches(x interface{}) bool {
	return m.x == x
}
