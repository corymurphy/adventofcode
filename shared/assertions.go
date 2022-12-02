package shared

import (
	"testing"
)

func AssertEqual(t *testing.T, expected interface{}, actual interface{}) {
	if expected != actual {
		t.Errorf("expected %s, got %s", expected, actual)
	}
}
