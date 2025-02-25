package matchers

import (
	"fmt"
	"go.uber.org/mock/gomock"
)

type SliceLengthMatcher[T any] struct {
	expectedLen int
}

func (s *SliceLengthMatcher[T]) Matches(x interface{}) bool {
	slice, ok := x.([]T)
	return ok && len(slice) == s.expectedLen
}

func (s *SliceLengthMatcher[T]) String() string {
	return fmt.Sprintf("%d", s.expectedLen)
}

func SliceLength[T any](expectedLen int) gomock.Matcher {
	return &SliceLengthMatcher[T]{expectedLen: expectedLen}
}
