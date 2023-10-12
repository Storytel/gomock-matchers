package matchers_test

import (
	matchers "github.com/Storytel/gomock-matchers"
	pb "github.com/Storytel/gomock-matchers/gomockmatchersproto"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProtoMatcher(t *testing.T) {
	assert := assert.New(t)
	p1 := &pb.TestRecord{Key: "a", Value: "1"}
	p2 := &pb.TestRecord{Key: "a", Value: "1"}
	m := matchers.Proto(p1)
	assert.True(m.Matches(p2))
}

func TestProtoMatcherString(t *testing.T) {
	assert := assert.New(t)
	p1 := &pb.TestRecord{Key: "a", Value: "1"}
	m := matchers.Proto(p1)
	assert.Equal("is matching", m.String())
}
