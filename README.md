# Gomock Matcher

Matching library for use with [golang/gomock][golang-gomock]

# Usage

All matchers in this package implements the [Matcher interface][matcher-interface] and can thusly be used with mocks from [golang/gomock][golang-gomock].

Example usage of the **RecordingMatcher**, but all matchers work in a similar wahy:

```go
import "github.com/Storytel/gomock-matchers"

func TestSomething(t *testing.T) {
  // Set up a new matcher
  matcher := matchers.NewRecordingMatcher(gomock.Eq(12))

  // Expect `MyFunc` to be called on `mock`
  mock.EXPECT().MyFunc(matcher).Return(nil)

  // Run the test this will expect MyFunc on Mock to be called with arg 12
  DUT()

  // For the RecordingMatcher, you can get the argument. Other matchers have other characteristics.
  if 12 != matcher.Get().(int) {
    t.Fail()
  }
}
```

# Matchers

<details>
<summary><strong>RecordingMatcher</strong> - <em>Proxy matcher which captures the argument for further inspection.</em></summary>
Wraps another matcher and records the value of the argument it's called with.

This can be used if you need to do further investigations. For instance when
the argument is a function, and you want to test that function.

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
</details>

[matcher-interface]: https://godoc.org/github.com/golang/mock/gomock#Matcher
[golang-gomock]: https://github.com/golang/mock