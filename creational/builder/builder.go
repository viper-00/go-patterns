package builder

import "fmt"

// wiki: https://en.wikipedia.org/wiki/Builder_pattern

/*
Builder is a creational design pattern that produce different types
and representations of an object using the same construction code.
*/

// builder interface
type Builder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() House
}

// concrete builder
type NormalBuilder struct {
	windoeType string
	doorType   string
	floor      int
}

func NewNormalBuilder() *NormalBuilder {
	return &NormalBuilder{}
}

func (normal *NormalBuilder) setWindowType() {
	normal.windoeType = "normal window"
}

func (normal *NormalBuilder) setDoorType() {
	normal.doorType = "normal door"
}

func (normal *NormalBuilder) setNumFloor() {
	normal.floor = 3
}

func (normal *NormalBuilder) getHouse() House {
	return House{
		windowType: normal.windoeType,
		doorType:   normal.doorType,
		floor:      normal.floor,
	}
}

// concrete builder
type IglooBuilder struct {
	windoeType string
	doorType   string
	floor      int
}

func NewIglooBuilder() *IglooBuilder {
	return &IglooBuilder{}
}

func (igloo *IglooBuilder) setWindowType() {
	igloo.windoeType = "igloo window"
}

func (igloo *IglooBuilder) setDoorType() {
	igloo.doorType = "igloo door"
}

func (igloo *IglooBuilder) setNumFloor() {
	igloo.floor = 1
}

func (igloo *IglooBuilder) getHouse() House {
	return House{
		windowType: igloo.windoeType,
		doorType:   igloo.doorType,
		floor:      igloo.floor,
	}
}

// product
type House struct {
	windowType string
	doorType   string
	floor      int
}

// director
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

func (d *Director) builderHouse() House {
	d.builder.setWindowType()
	d.builder.setDoorType()
	d.builder.setNumFloor()
	return d.builder.getHouse()
}

// get a builder
func GetBuilder(builderType string) (Builder, error) {
	if builderType == "normal" {
		return &NormalBuilder{}, nil
	} else if builderType == "igloo" {
		return &IglooBuilder{}, nil
	}

	return nil, fmt.Errorf("wrong builder type passed")
}
