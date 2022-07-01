package assert

import (
	"fmt"
	"strings"
)

type T interface {
	Fatalf(string, ...interface{})
	Helper()
	Name() string
}

func message0(t T, message string, messages ...string) {
	t.Fatalf("%s failed, %s:\n\n%s", t.Name(), message, strings.Join(messages, ", "))
}

func message1(t T, message string, left interface{}, messages ...string) {
	message0(t, message, fmt.Sprintf("got: %v\n\n", left))
}

func message2(t T, message string, left interface{}, right interface{}, messages ...string) {
	message0(t, message, fmt.Sprintf("left:\n  %v\n\nright:\n  %v\n\n", left, right)) //, messages...)
}

func Equal(t T, left interface{}, right interface{}, messages ...string) {
	t.Helper()
	if left != right {
		message2(t, "should be equal", left, right, messages...)
	}
}

func NotEqual(t T, left interface{}, right interface{}, messages ...string) {
	t.Helper()
	if left == right {
		message2(t, "should not be equal", left, right, messages...)
	}
}

func Contains(t T, left string, right string) {
	t.Helper()

	if !strings.Contains(left, right) {
		message2(t, "should contain substring", left, right)
	}
}

func Excludes(t T, left string, right string) {
	t.Helper()

	if strings.Contains(left, right) {
		message2(t, "should not contain substring", left, right)
	}
}

func True(t T, got bool) {
	t.Helper()
	if !got {
		message1(t, "should be true", got)
	}
}

func False(t T, got bool) {
	t.Helper()
	if got {
		message1(t, "should be false", got)
	}
}

func Nil(t T, got interface{}) {
	t.Helper()
	if got != nil {
		message1(t, "should be nil", got)
	}
}

func NotNil(t T, got interface{}) {
	t.Helper()
	if got == nil {
		message0(t, "should not be nil")
	}
}

func Error(t T, got error) {
	t.Helper()
	if got == nil {
		message1(t, "should throw an error", got)
	}
}

func Success(t T, got error) {
	t.Helper()
	if got != nil {
		message1(t, "should not throw an error", got)
	}
}
