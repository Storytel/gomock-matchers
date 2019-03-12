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

func NewRecordingMatcher(m gomock.Matcher) *recordingMatcher {
	return &recordingMatcher{
		m: m,
	}
}
