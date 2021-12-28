package flyweight

import "fmt"

// wiki: https://en.wikipedia.org/wiki/Flyweight_pattern
//
// Reuses existing instances of objects with similar/identical state to minimize resource usage.

const (
	TerroristDressType        = "tDress"
	CounterTerroristDressType = "ctDress"
)

type dressFactory struct {
	dressMap map[string]dress
}

var df dressFactory = dressFactory{
	dressMap: make(map[string]dress),
}

func GetDressFactorySingleInstance() *dressFactory {
	return &df
}

func (d *dressFactory) getDressByType(dressType string) (dress, error) {
	if dressValue, ok := d.dressMap[dressType]; ok {
		return dressValue, nil
	}

	switch dressType {
	case TerroristDressType:
		d.dressMap[dressType] = newTerroristDress()
		return d.dressMap[dressType], nil
	case CounterTerroristDressType:
		d.dressMap[dressType] = newCounterTerroristDress()
		return d.dressMap[dressType], nil
	default:
		return nil, fmt.Errorf("input dressType is invalid")
	}
}

type dress interface {
	getColor() string
}

type terroristDress struct {
	color string
}

func newTerroristDress() dress {
	return &terroristDress{color: "red"}
}

func (t *terroristDress) getColor() string {
	return t.color
}

type counterTerroristDress struct {
	color string
}

func newCounterTerroristDress() dress {
	return &counterTerroristDress{color: "green"}
}

func (c *counterTerroristDress) getColor() string {
	return c.color
}

type player struct {
	dress      dress
	playerType string
	lat        int
	long       int
}

func newPlayer(playerType string, dressType string) *player {
	dress, err := GetDressFactorySingleInstance().getDressByType(dressType)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &player{
		playerType: playerType,
		dress:      dress,
	}
}

func (p *player) newLocation(lat, long int) {
	p.lat = lat
	p.long = long
}

type game struct {
	terrorists        []*player
	counterTerrorists []*player
}

func NewGame() *game {
	return &game{}
}

var (
	terroristplayerType        = "T"
	counterTerroristplayerType = "CT"
)

func (g *game) addTerrorist(dressType string) {
	if len(g.terrorists) > 1 {
		return
	}
	player := newPlayer(terroristplayerType, dressType)
	player.newLocation(0, 0)
	g.terrorists = append(g.terrorists, player)
}

func (g *game) addCounterTerrorist(dressType string) {
	if len(g.counterTerrorists) > 1 {
		return
	}
	player := newPlayer(counterTerroristplayerType, dressType)
	player.newLocation(10, 10)
	g.counterTerrorists = append(g.counterTerrorists, player)
}
