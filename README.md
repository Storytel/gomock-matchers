# Gomock Matcher

Matching library for use with [golang/gomock][golang-gomock]

## Available matchers

  * [RecordingMatcher](#recording-matcher) - Proxy matcher which captures the argument.

## Usage

All matchers in this package implements the [Matcher interface][matcher-interface] and can thusly be used with mocks from [golang/gomock][golang-gomock].
Example usage of the RecordingMatcher, but all matchers work in a similar wahy:

```go
matcher := 
mock.EXPECT().Function().Return(nil)
```

## Recording Matcher

Wraps another matcher and records the value of the argument it's called with.
This can be used if you need to do further investigations. For instance when
the argument is a function, and you want to test that function you need a way
to get to it.

```go
type MyFunc func() int

func TestSomething (t *testing.T) {
	assert := assert.New(t)
	m := NewRecordingMatcher(gomock.Any())

	m.Matches(MyFunc(func () int { return 12 }))

	f, ok := m.Get().(MyFunc)
	assert.True(ok)
	assert.Equal(12, f())
}
```

[matcher-interface]: https://godoc.org/github.com/golang/mock/gomock#Matcher
[golang-gomock]: https://github.com/golang/mock