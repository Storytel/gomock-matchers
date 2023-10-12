package matchers

import (
	"github.com/golang/mock/gomock"
	"google.golang.org/protobuf/proto"
)

type ProtoMatcher struct {
	data proto.Message
}

func Proto(x proto.Message) gomock.Matcher {
	return &ProtoMatcher{data: x}
}

func (p *ProtoMatcher) Matches(x interface{}) bool {
	pbx, ok := x.(proto.Message)
	if !ok {
		return false
	}
	return proto.Equal(p.data, pbx)
}

func (p *ProtoMatcher) String() string {
	return "is matching"
}
