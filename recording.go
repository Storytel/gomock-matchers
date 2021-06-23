package matchers

import "github.com/golang/mock/gomock"

type RecordMatcher struct {
	x interface{}
	m gomock.Matcher
}

// Record returns a new matcher which wraps the
// provided matcher - following all the matching rules of that
// matcher. In addition, the argument which is matched is recorded
// and can later be retrieved for inspection using the Get() func
func Record(m gomock.Matcher) *RecordMatcher {
	return &RecordMatcher{
		m: m,
	}
}

func (rm *RecordMatcher) Matches(x interface{}) bool {
	rm.x = x
	return rm.m.Matches(x)
}

func (rm *RecordMatcher) String() string {
	return rm.m.String()
}

func (rm *RecordMatcher) Get() interface{} {
	return rm.x
}
