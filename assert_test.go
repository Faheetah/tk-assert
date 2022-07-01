package assert_test

import (
	"errors"
	"fmt"
	"testing"

	assert "github.com/faheetah/tk-assert"
)

type fakeT struct {
	Failed  bool
	Message string
}

func (f *fakeT) Fatalf(message string, args ...interface{}) {
	f.Failed = true
	f.Message = fmt.Sprintf(message, args...)
}

func (f *fakeT) Helper() {}

func (f *fakeT) Name() string { return "TestExample" }

func TestEqual(t *testing.T) {
	mockSuccess := &fakeT{}
	assert.Equal(mockSuccess, 1, 1)
	assert.False(t, mockSuccess.Failed)

	mockFailed := &fakeT{}
	assert.Equal(mockFailed, 1, 2)
	assert.Contains(t, mockFailed.Message, "should be equal")
	assert.Contains(t, mockFailed.Message, "\n\nleft:\n  1")
	assert.Contains(t, mockFailed.Message, "\n\nright:\n  2")
}

func TestNotEqual(t *testing.T) {
	mockSuccess := &fakeT{}
	assert.NotEqual(mockSuccess, 1, 2)
	assert.False(t, mockSuccess.Failed)

	mockFailed := &fakeT{}
	assert.NotEqual(mockFailed, 1, 1)
	assert.Contains(t, mockFailed.Message, "should not be equal")
}

func TestContains(t *testing.T) {
	mockSuccess := &fakeT{}
	assert.Contains(mockSuccess, "foo", "foo")
	assert.False(t, mockSuccess.Failed)

	mockFailed := &fakeT{}
	assert.Contains(mockFailed, "foo", "bar")
	assert.Contains(t, mockFailed.Message, "should contain")
}

func TestExcludes(t *testing.T) {
	mockSuccess := &fakeT{}
	assert.Excludes(mockSuccess, "foo", "bar")
	assert.False(t, mockSuccess.Failed)

	mockFailed := &fakeT{}
	assert.Excludes(mockFailed, "foo", "foo")
	assert.Contains(t, mockFailed.Message, "should not contain")
}

func TestTrue(t *testing.T) {
	mockSuccess := &fakeT{}
	assert.True(mockSuccess, true)
	assert.False(t, mockSuccess.Failed)

	mockFailed := &fakeT{}
	assert.True(mockFailed, false)
	assert.Contains(t, mockFailed.Message, "should be true")
	assert.Contains(t, mockFailed.Message, "got: false")
}

func TestFalse(t *testing.T) {
	mockSuccess := &fakeT{}
	assert.False(mockSuccess, false)
	assert.False(t, mockSuccess.Failed)

	mockFailed := &fakeT{}
	assert.False(mockFailed, true)
	assert.Contains(t, mockFailed.Message, "should be false")
	assert.Contains(t, mockFailed.Message, "got: true")
}

func TestNil(t *testing.T) {
	mockSuccess := &fakeT{}
	assert.Nil(mockSuccess, nil)
	assert.False(t, mockSuccess.Failed)

	mockFailed := &fakeT{}
	assert.Nil(mockFailed, 1)
	assert.Contains(t, mockFailed.Message, "should be nil")
	assert.Contains(t, mockFailed.Message, "got")
}

func TestNotNil(t *testing.T) {
	mockSuccess := &fakeT{}
	assert.NotNil(mockSuccess, 1)
	assert.False(t, mockSuccess.Failed)

	mockFailed := &fakeT{}
	assert.NotNil(mockFailed, nil)
	assert.Contains(t, mockFailed.Message, "should not be nil")
}

func TestError(t *testing.T) {
	mockSuccess := &fakeT{}
	assert.Error(mockSuccess, errors.New("failed"))
	assert.False(t, mockSuccess.Failed)

	mockFailed := &fakeT{}
	assert.Error(mockFailed, nil)
	assert.Contains(t, mockFailed.Message, "should throw an error")
}

func TestSuccess(t *testing.T) {
	mockSuccess := &fakeT{}
	assert.Success(mockSuccess, nil)
	assert.False(t, mockSuccess.Failed)

	mockFailed := &fakeT{}
	assert.Success(mockFailed, errors.New("failed"))
	assert.Contains(t, mockFailed.Message, "should not throw an error")
}
