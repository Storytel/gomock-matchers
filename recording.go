package matchers

import "github.com/golang/mock/gomock"

type recordMatcher struct {
	x interface{}
	m gomock.Matcher
}

// Record returns a new matcher which wraps the
// provided matcher - following all the matching rules of that
// matcher. In addition, the argument which is matched is recorded
// and can later be retrieved for inspection using the Get() func
func Record(m gomock.Matcher) *recordMatcher {
	return &recordMatcher{
		m: m,
	}
}

func (rm *recordMatcher) Matches(x interface{}) bool {
	rm.x = x
	return rm.m.Matches(x)
}

func (rm *recordMatcher) String() string {
	return rm.m.String()
}

func (rm *recordMatcher) Get() interface{} {
	return rm.x
}
