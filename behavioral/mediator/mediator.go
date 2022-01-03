package mediator

import "fmt"

// wiki: https://en.wikipedia.org/wiki/Mediator_pattern
//
// Connects objects and acts as a proxy.

type train interface {
	arrive()
	depart()
	permitArrival()
}

type mediator interface {
	canArrive(train) bool
	notifyAboutDeparture()
}

type passengerTrain struct {
	mediator mediator
}

func (p *passengerTrain) arrive() {
	if !p.mediator.canArrive(p) {
		fmt.Println("PassengerTrain: Arrival blocked, waiting")
		return
	}

	fmt.Println("PassengerTrain: Arriving")
}

func (p *passengerTrain) depart() {
	fmt.Println("PassengerTrain: Leaving")
	p.mediator.notifyAboutDeparture()
}

func (p *passengerTrain) permitArrival() {
	fmt.Println("PassengerTrain: Arrival permitted, arriving")
	p.arrive()
}

type freightTrain struct {
	mediator mediator
}

func (f *freightTrain) arrive() {
	if !f.mediator.canArrive(f) {
		fmt.Println("FreightTrain: Arrival blocked, waiting")
		return
	}
	fmt.Println("FreightTrain: Arriving")
}

func (f *freightTrain) depart() {
	fmt.Println("FreightTrain: Leaving")
	f.mediator.notifyAboutDeparture()
}

func (f *freightTrain) permitArrival() {
	fmt.Println("FreightTrain: Arrival permitted")
	f.arrive()
}

type stationManager struct {
	isPlatformFree bool
	trainQueue     []train
}

func newStationManager() *stationManager {
	return &stationManager{
		isPlatformFree: true,
	}
}

func (s *stationManager) canArrive(t train) bool {
	if s.isPlatformFree {
		s.isPlatformFree = false
		return true
	}
	s.trainQueue = append(s.trainQueue, t)
	return false
}

func (s *stationManager) notifyAboutDeparture() {
	if !s.isPlatformFree {
		s.isPlatformFree = true
	}
	if len(s.trainQueue) > 0 {
		firstTrainInQueue := s.trainQueue[0]
		s.trainQueue = s.trainQueue[1:]
		firstTrainInQueue.permitArrival()
	}
}
