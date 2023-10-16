package matchers_test

import (
	"testing"

	matchers "github.com/Storytel/gomock-matchers/v2"
	"github.com/stretchr/testify/assert"

	"google.golang.org/protobuf/types/known/structpb"
)

func TestProtoMatcher(t *testing.T) {
	data := &structpb.Struct{
		Fields: map[string]*structpb.Value{
			"Id": {
				Kind: &structpb.Value_StringValue{
					StringValue: "1234",
				},
			},
		},
	}

	assert := assert.New(t)
	m := matchers.Proto(data)
	assert.True(m.Matches(data))
}

func TestProtoMatcherString(t *testing.T) {
	assert := assert.New(t)
	data := &structpb.Struct{
		Fields: map[string]*structpb.Value{
			"abc": {
				Kind: &structpb.Value_StringValue{
					StringValue: "123",
				},
			},
		},
	}

	expected := "{\"abc\":\"123\"}"
	m := matchers.Proto(data)
	assert.Equal(expected, m.String())
}
