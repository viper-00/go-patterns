package mediator

import "testing"

func TestPattern(t *testing.T) {
	stationManager := newStationManager()

	passengerTrain := &passengerTrain{mediator: stationManager}

	freightTrain := &freightTrain{mediator: stationManager}

	passengerTrain.arrive()
	freightTrain.arrive()
	passengerTrain.depart()
}
