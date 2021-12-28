package flyweight

import "testing"

func TestPattern(t *testing.T) {
	game := NewGame()

	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)

	game.addCounterTerrorist(CounterTerroristDressType)
	game.addCounterTerrorist(CounterTerroristDressType)
	game.addCounterTerrorist(CounterTerroristDressType)

	dressFactory := GetDressFactorySingleInstance()

	for dressType, dress := range dressFactory.dressMap {
		t.Logf("Dress Type: %s DressColor: %s\n", dressType, dress.getColor())
	}
}
