package builder

import "fmt"

// wiki: https://en.wikipedia.org/wiki/Builder_pattern
//
// Builds a complex object using simple objects.

type bType int

const (
	normal bType = 1
	igloo  bType = 2
)

type Builder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() *House
}

type NormalBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func NewNormalBuilder() *NormalBuilder {
	return new(NormalBuilder)
}

func (normal *NormalBuilder) setWindowType() {
	normal.windowType = "normal window"
}

func (normal *NormalBuilder) setDoorType() {
	normal.doorType = "normal door"
}

func (normal *NormalBuilder) setNumFloor() {
	normal.floor = 2
}

func (normal *NormalBuilder) getHouse() *House {
	return &House{
		windowType: normal.windowType,
		doorType:   normal.doorType,
		floor:      normal.floor,
	}
}

type IglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func NewIglooBuilder() *IglooBuilder {
	return new(IglooBuilder)
}

func (igloo *IglooBuilder) setWindowType() {
	igloo.windowType = "igloo window"
}

func (igloo *IglooBuilder) setDoorType() {
	igloo.doorType = "igloo door"
}

func (igloo *IglooBuilder) setNumFloor() {
	igloo.floor = 1
}

func (igloo *IglooBuilder) getHouse() *House {
	return &House{
		windowType: igloo.windowType,
		doorType:   igloo.doorType,
		floor:      igloo.floor,
	}
}

type House struct {
	windowType string
	doorType   string
	floor      int
}

type Director struct {
	builder Builder
}

func NewDirector(b Builder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) setBuilder(b Builder) {
	d.builder = b
}

func (d *Director) builderHouse() *House {
	d.builder.setWindowType()
	d.builder.setDoorType()
	d.builder.setNumFloor()
	return d.builder.getHouse()
}

func GetBuilder(typ bType) (Builder, error) {
	switch typ {
	case normal:
		return new(NormalBuilder), nil
	case igloo:
		return new(IglooBuilder), nil
	default:
		return nil, fmt.Errorf(fmt.Sprintf("Builder of type %d is not exist", typ))
	}
}
