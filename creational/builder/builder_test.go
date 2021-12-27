package builder

import (
	"strconv"
	"testing"
)

func TestPattern(t *testing.T) {
	normalBuilder, err1 := GetBuilder(normal)

	if err1 != nil {
		t.Error(err1)
	}

	iglooBuilder, err2 := GetBuilder(igloo)
	if err2 != nil {
		t.Error(err2)
	}

	director := NewDirector(normalBuilder)
	normalHouse := director.builderHouse()

	t.Log("Normal House Window Type: " + normalHouse.windowType)
	t.Log("Normal House Door Type: " + normalHouse.doorType)
	t.Log("Normal House Floor: " + strconv.Itoa(normalHouse.floor))

	director.setBuilder(iglooBuilder)
	iglooHouse := director.builderHouse()

	t.Log("Igloo House Window Type: " + iglooHouse.windowType)
	t.Log("Igloo House Door Type: " + iglooHouse.doorType)
	t.Log("Igloo House Floor: " + strconv.Itoa(iglooHouse.floor))
}
