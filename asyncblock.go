package matchers

import "github.com/golang/mock/gomock"

type asyncBlockMatcher struct {
	matcher gomock.Matcher
	ch      chan struct{}
}

// AsyncBlock returns a matcher holding a channel which will be signaled when
// the `Matches` function is called. AsyncBlock will wrap any other matcher
// to do the actual matching.
//
// This is useful if the code you're testing is spawning a go function which will
// invoke your mock at some time in the future. The channel gives you an easy way
// to wait for that invokation (using `<- matcher.Channel()`) and then do assertions.
func AsyncBlock(matcher gomock.Matcher) *asyncBlockMatcher {
	return &asyncBlockMatcher{
		ch:      make(chan struct{}),
		matcher: matcher,
	}
}

// Channel returns a channel which will have an empty struct written to it every time
// `Matches` is invoked.
func (m *asyncBlockMatcher) Channel() <-chan struct{} {
	return m.ch
}

func (m *asyncBlockMatcher) Matches(x interface{}) bool {
	b := m.matcher.Matches(x)
	m.ch <- struct{}{}
	return b
}

func (m *asyncBlockMatcher) String() string {
	return m.matcher.String()
}
