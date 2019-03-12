package matchers

import "github.com/golang/mock/gomock"

type recordingMatcher struct {
	x interface{}
	m gomock.Matcher
}

func (rm *recordingMatcher) Matches(x interface{}) bool {
	rm.x = x
	return rm.m.Matches(x)
}

func (rm *recordingMatcher) String() string {
	return rm.m.String()
}

func (rm *recordingMatcher) Get() interface{} {
	return rm.x
}

// NewRecordingMatcher returns a new matcher which wraps the
// provided matcher - following all the matching rules of that
// matcher. In addition, the argument which is matched is recorded
// and can later be retrieved for inspection using the Get() func
func NewRecordingMatcher(m gomock.Matcher) *recordingMatcher {
	return &recordingMatcher{
		m: m,
	}
}
