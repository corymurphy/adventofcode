package shared

import (
	"testing"
)
// import (
// 	"ioutil"
// )

// func AssertFileContent(t *testing.T, expected string, actualPath string) {
// 	actual, err := ioutil.ReadFile(actualPath)

// 	if err != nil {
// 		t.Errorf("actual file %s not found", actualPath)
// 	}

// 	if expected != string(actual) {
// 		t.Errorf("expected %s, got %s", expected, string(actual))
// 	}
// }

func AssertEqual(t *testing.T, expected interface{}, actual interface{}) {
	if expected != actual {
		t.Errorf("expected %s, got %s", expected, actual)
	}
}

// func AssertIntEqual(t *testing.T, expected int, actual int) {
// 	if expected != actual {
// 		t.Errorf("expected %s, got %s", expected, actual)
// 	}
// }

