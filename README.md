# Gomock Matcher

[![CircleCI](https://circleci.com/gh/Storytel/gomock-matchers/tree/master.svg?style=svg)](https://circleci.com/gh/Storytel/gomock-matchers/tree/master)
[![codecov](https://codecov.io/gh/Storytel/gomock-matchers/branch/master/graph/badge.svg)](https://codecov.io/gh/Storytel/gomock-matchers)
[![Go Report Card](https://goreportcard.com/badge/github.com/Storytel/gomock-matchers)](https://goreportcard.com/report/github.com/Storytel/gomock-matchers)

Matching library for use with [golang/gomock][golang-gomock]

# Usage

All matchers in this package implements the [Matcher interface][matcher-interface] and can thusly be used with mocks from [golang/gomock][golang-gomock].

Example usage of **Record** matcher, but all matchers work in a similar way:

```go
import "github.com/Storytel/gomock-matchers"

func TestSomething(t *testing.T) {
  // Set up a new matcher
  matcher := matchers.Record(gomock.Eq(12))

  // Expect `MyFunc` to be called on `mock` with argument equal to 12
  mock.EXPECT().MyFunc(matcher).Return(nil)

  // Run the test this which will expect the EXPECT()
  DUT()

  // For the RecordMatcher, you can get the argument. Other matchers have other characteristics.
  if 12 != matcher.Get().(int) {
    t.Fail()
  }
}
```

# Matchers

<details>
<summary><strong>RecordMatcher</strong> - <em>Proxy matcher which captures the argument for further inspection.</em></summary>
Wraps another matcher and records the value of the argument it's called with.

This can be used if you need to do further investigations. For instance when
the argument is a function, and you want to test that function.

```go
type MyFunc func() int

func TestRecord(t *testing.T) {
	assert := assert.New(t)
	m := matchers.Record(gomock.Any())

	m.Matches(MyFunc(func () int { return 12 }))

	f, ok := m.Get().(MyFunc)
	assert.True(ok)
	assert.Equal(12, f())
}
```
</details>

<details>
<summary><strong>SameMatcher</strong> - <em>Matcher which checks if values are the same</em></summary>

This differs from `gomock.Eq` in that it does a comparison check with `==` and not a
`reflect.DeepEqual`. This means that two pointers are only _same_ if they point to the
same memory address

```go
func TestSame(t *testing.T) {
  assert := assert.New(t)

  myString := "something"
  otherString := "something"

  m := matchers.Same(&myString)
  assert.True(m.Matches(&myString))
  assert.False(m.Matches(&otherString)) // Not the same pointer

  m2 := matchers.Same(myString)
  assert.True(m2.Matches(myString))
  assert.True(m2.Matches(otherString)) // Not pointers, values are the same
}
```
</details>

<summary><strong>AsyncBlockMatcher</strong> - <em>Matcher which provides channel signaling when `Matches` is called</em></summary>

AsyncBlock returns a matcher holding a channel which will be signaled when
the `Matches` function is called. AsyncBlock will wrap any other matcher
to do the actual matching.

This is useful if the code you're testing is spawning a go function which will
invoke your mock at some time in the future. The channel gives you an easy way
to wait for that invokation (using `<- matcher.Channel()`) and then do assertions.

```go
func TestAsyncBlockMatcher(t *testing.T) {
	assert := assert.New(t)
	m := matchers.AsyncBlock(gomock.Eq("12"))

	didMatch := false
	go func() {
		didMatch = m.Matches("12")
	}()

  // This blocks until `Matches` is actually called
	<-m.Channel()
  assert.True(didMatch)
}
```
</details>

[matcher-interface]: https://godoc.org/github.com/golang/mock/gomock#Matcher
[golang-gomock]: https://github.com/golang/mock
