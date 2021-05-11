package flyweight

import "testing"

func TestPattern(t *testing.T) {
	game := newGame()

	// add terrorist
	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)

	// add counterTerrorist
	game.addCounterTerrorist(CounterTerroristDressType)
	game.addCounterTerrorist(CounterTerroristDressType)
	game.addCounterTerrorist(CounterTerroristDressType)

	dressFactory := getDressFactorySingleInstance()

	for dressType, dress := range dressFactory.dressMap {
		t.Logf("Dress Type: %s DressColor: %s\n", dressType, dress.getColor())
	}
}
