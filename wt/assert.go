package wt

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func ExpectError(t *testing.T, err error, format string, args ...interface{}) {
	t.Helper()
	if err == nil {
		t.Errorf("Expected an error, but got nil [%s]", fmt.Sprintf(format, args...))
	}
}

func ExpectNoError(t *testing.T, err error, format string, args ...interface{}) {
	t.Helper()
	if err != nil {
		t.Errorf("Expected no error, but got %v [%s]", err, fmt.Sprintf(format, args...))
	}
}

func ExpectEqual(t *testing.T, expected, actual interface{}) {
	t.Helper()
	diff := cmp.Diff(expected, actual)
	if diff != "" {
		t.Errorf("invalid lajkoHTML rendered: %s", diff)
	}
}

func ExpectNil(t *testing.T, actual interface{}, format string, args ...interface{}) {
	t.Helper()
	if actual != nil {
		t.Errorf("Expected nil, but got %+#v [%s]", actual, fmt.Sprintf(format, args...))
	}
}
