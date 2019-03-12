package matchers_test

import (
	"testing"

	. "github.com/Storytel/gomock-matchers"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type MyType int

func TestRecordingMatcherPassthrough(t *testing.T) {
	assert := assert.New(t)
	m := NewRecordingMatcher(gomock.Eq(MyType(12)))

	assert.False(m.Matches(12))
	assert.True(m.Matches(MyType(12)))
}

func TestRecordingMatcherGetValue(t *testing.T) {
	assert := assert.New(t)
	m := NewRecordingMatcher(gomock.Any())

	assert.True(m.Matches("this is my thing"))
	assert.Equal("this is my thing", m.Get())
}

func TestRecordingMatcherString(t *testing.T) {
	assert := assert.New(t)
	m := NewRecordingMatcher(gomock.Any())
	assert.Equal("is anything", m.String()) // This is String() from gomock.anyMatcher
}
