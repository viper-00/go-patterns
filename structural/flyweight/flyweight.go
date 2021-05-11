package flyweight

import "fmt"

// wiki: https://en.wikipedia.org/wiki/Flyweight_pattern

/**
 * Flyweight is a structural design pattern that fit more objects into the available
 * amount of RAM by sharing comomon parts of state between multiple objects
 * instead of keeping all of data in each object.
 */

// flyweight factory
const (
	TerroristDressType        = "tDress"
	CounterTerroristDressType = "ctDress"
)

var (
	dressFactorySingleInstance = &dressFactory{
		dressMap: make(map[string]dress),
	}
)

type dressFactory struct {
	dressMap map[string]dress
}

func (d *dressFactory) getDressByType(dressType string) (dress, error) {
	if d.dressMap[dressType] != nil {
		return d.dressMap[dressType], nil
	}

	if dressType == TerroristDressType {
		d.dressMap[dressType] = newTerroristDress()
		return d.dressMap[dressType], nil
	} else if dressType == CounterTerroristDressType {
		d.dressMap[dressType] = newCounterTerroristDress()
		return d.dressMap[dressType], nil
	}

	return nil, fmt.Errorf("wrong dress type passed")
}

func getDressFactorySingleInstance() *dressFactory {
	return dressFactorySingleInstance
}

// flyweight interface
type dress interface {
	getColor() string
}

// concrete flyweight object - terroristDress
type terroristDress struct {
	color string
}

func (t *terroristDress) getColor() string {
	return t.color
}

func newTerroristDress() dress {
	return &terroristDress{color: "red"}
}

// concrete flyweight object - counterTerroristDress
type counterTerroristDress struct {
	color string
}

func (c *counterTerroristDress) getColor() string {
	return c.color
}

func newCounterTerroristDress() dress {
	return &counterTerroristDress{color: "green"}
}

// context - player
type player struct {
	dress      dress
	playerType string
	lat        int
	long       int
}

func newPlayer(playerType string, dressType string) *player {
	dress, _ := getDressFactorySingleInstance().getDressByType(dressType)
	return &player{
		playerType: playerType,
		dress:      dress,
	}
}

func (p *player) newLocation(lat, long int) {
	p.lat = lat
	p.long = long
}

// client code - game
type game struct {
	terrorists        []*player
	counterTerrorists []*player
}

func newGame() *game {
	return &game{
		terrorists:        make([]*player, 1),
		counterTerrorists: make([]*player, 1),
	}
}

func (g *game) addTerrorist(dressType string) {
	player := newPlayer("T", dressType)
	player.newLocation(0, 0)
	g.terrorists = append(g.terrorists, player)
}

func (g *game) addCounterTerrorist(dressType string) {
	player := newPlayer("CT", dressType)
	player.newLocation(10, 10)
	g.counterTerrorists = append(g.counterTerrorists, player)
}
