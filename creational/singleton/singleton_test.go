package singleton

import (
	"testing"
)

func TestPattern(t *testing.T) {

	for i := 0; i < 30; i++ {
		go GetInstance()
	}
}
