package matchers_test

import (
	"errors"
	"fmt"
	"testing"

	matchers "github.com/Storytel/gomock-matchers/v2"

	"github.com/stretchr/testify/assert"
)

func TestSameMatcher(t *testing.T) {

	str1 := "something"
	str2 := "something"
	err1 := errors.New("error")
	err2 := errors.New("error")

	tests := []struct {
		lhs  interface{}
		rhs  interface{}
		same bool
	}{
		{"asdf", "asdf", true},
		{&err1, &err1, true},
		{12, 12, true},
		{&str1, &str1, true},

		{&str1, &str2, false},
		{&err1, &err2, false},
		{"bsdf", "asdf", false},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Test #%d", i), func(t *testing.T) {
			assert := assert.New(t)
			m := matchers.Same(test.lhs)
			asserter := assert.False
			if test.same {
				asserter = assert.True
			}
			asserter(m.Matches(test.rhs))
		})
	}
}

type StringTestType int

func (s StringTestType) String() string {
	return "StringTestType#String()"
}

func TestSameMatcherString(t *testing.T) {
	tests := []struct {
		v interface{}
		s string
	}{
		{"asdf", "is same as 'asdf'"},
		{12, "is same as '12'"},
		{StringTestType(12), "is same as 'StringTestType#String()'"},
		{[]int{1, 2, 3}, "is same as '[1 2 3]'"},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Test #%d", i), func(t *testing.T) {
			assert := assert.New(t)
			m := matchers.Same(test.v)
			assert.Equal(test.s, m.String())
		})
	}
}

func TestSame(t *testing.T) {
	assert := assert.New(t)

	myString := "something"
	otherString := "something"

	m := matchers.Same(&myString)
	assert.True(m.Matches(&myString))
	assert.False(m.Matches(&otherString)) // Not the same pointer

	m2 := matchers.Same(myString)
	assert.True(m2.Matches(myString))
	assert.True(m2.Matches(otherString)) // Same value because they're not pointers
}
