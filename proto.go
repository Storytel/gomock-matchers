package matchers

import (
	"encoding/json"

	"github.com/golang/mock/gomock"
	"google.golang.org/protobuf/proto"
)

type ProtoMatcher struct {
	data proto.Message
}

func (p *ProtoMatcher) Matches(x interface{}) bool {
	pbx, ok := x.(proto.Message)
	if !ok {
		return false
	}
	return proto.Equal(p.data, pbx)
}

func (p *ProtoMatcher) String() string {
	data, err := json.Marshal(p.data)
	if err != nil {
		return "<unknown>"
	}
	return string(data)
}

func Proto(x proto.Message) gomock.Matcher {
	return &ProtoMatcher{data: x}
}
