package matchers

import "fmt"

type sameMatcher struct {
	x interface{}
}

func Same(x interface{}) *sameMatcher {
	return &sameMatcher{x}
}

func (m *sameMatcher) String() string {
	return fmt.Sprintf("is same as '%v'", m.x)
}

func (m *sameMatcher) Matches(x interface{}) bool {
	return m.x == x
}
