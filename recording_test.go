package matchers_test

import (
	"testing"

	matchers "github.com/Storytel/gomock-matchers/v2"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestRecordMatcherPassthrough(t *testing.T) {
	type MyType int

	assert := assert.New(t)
	m := matchers.Record(gomock.Eq(MyType(12)))

	assert.False(m.Matches(12))
	assert.True(m.Matches(MyType(12)))
}

func TestRecordMatcherGetValue(t *testing.T) {
	assert := assert.New(t)
	m := matchers.Record(gomock.Any())

	assert.True(m.Matches("this is my thing"))
	assert.Equal("this is my thing", m.Get())
}

func TestRecordMatcherString(t *testing.T) {
	assert := assert.New(t)
	m := matchers.Record(gomock.Any())
	assert.Equal("is anything", m.String()) // This is String() from gomock.anyMatcher
}
