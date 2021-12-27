package singleton

import (
	"testing"
)

func TestPattern(t *testing.T) {
	s1 := New()
	var expectedKey, expectedValue string = "this", "that"
	s1[expectedKey] = expectedValue
	s2 := New()
	if actualValue := s2[expectedKey]; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
}
