package matchers_test

import (
	"testing"

	matchers "github.com/Storytel/gomock-matchers"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestAsyncBlockMatcherPassthrough(t *testing.T) {
	type MyType int

	assert := assert.New(t)
	m := matchers.AsyncBlock(gomock.Eq(MyType(24)))

	go func() { <-m.Channel() }()
	assert.False(m.Matches(24))

	go func() { <-m.Channel() }()
	assert.True(m.Matches(MyType(24)))
}

func TestAsyncBlockMatcher(t *testing.T) {

	assert := assert.New(t)
	m := matchers.AsyncBlock(gomock.Eq("12"))

	ch := make(chan bool)
	go func() {
		ch <- m.Matches("12")
	}()

	<-m.Channel()
	assert.True(<-ch)
}

func TestAsyncBlockMatcherString(t *testing.T) {
	assert := assert.New(t)
	m := matchers.AsyncBlock(gomock.Any())
	assert.Equal("is anything", m.String()) // This is String() from gomock.anyMatcher
}
