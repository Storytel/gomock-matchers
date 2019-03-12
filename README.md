# Gomock Matcher

Matching library for use with [golang/gomock](https://github.com/golang/mock)

## Available matchers

  * [RecordingMatcher](recording-matcher) - Proxy matcher which captures the argument.


## Recording Matcher

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