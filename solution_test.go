package solution

import (
	"testing"
)

func TestGetMessage(t *testing.T) {
	expected := "Hello 🗺️ !"
	result := GetMessage()
	if result == "" {
		t.Error("message must not be empty")
	}
	if result != expected {
		t.Errorf("message content is wrong, expect:%s but was %s", expected, result)
	}
}
