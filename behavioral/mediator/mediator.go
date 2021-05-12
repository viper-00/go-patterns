package mediator

import "fmt"

// wiki: https://en.wikipedia.org/wiki/Mediator_pattern

/**
 * Mediator is a behavioral design pattern that reduces coupling between components of a program by making
 * them communicate indirectly through a special madiator object.
 */

// component - train
type train interface {
	arrive()
	depart()
	permitArrival()
}

// concrete component - passengerTrain
type passengerTrain struct {
	mediator mediator
}

func (p *passengerTrain) arrive() {
	if !p.mediator.canArrive(p) {
		fmt.Println("PassengerTrain: Arrival blocked, waiting")
		return
	}

	fmt.Println("PassengerTrain: Arrived")
}

func (p *passengerTrain) depart() {
	fmt.Println("PassengerTrain: Leaving")
	p.mediator.notifyAboutDeparture()
}

func (p *passengerTrain) permitArrival() {
	fmt.Println("PassengerTrain: Arrival permitted, arriving")
	p.arrive()
}

// concrete component - freightTrain
type freightTrain struct {
	mediator mediator
}

func (f *freightTrain) arrive() {
	if !f.mediator.canArrive(f) {
		fmt.Println("FreightTrain: Arrival blocked, waiting")
		return
	}
	fmt.Println("FreightTrain: Arrived")
}

func (f *freightTrain) depart() {
	fmt.Println("FreightTrain: Leaving")
	f.mediator.notifyAboutDeparture()
}

func (f *freightTrain) permitArrival() {
	fmt.Println("FreightTrain: Arrival permitted")
	f.arrive()
}

// mediator interface - mediator
type mediator interface {
	canArrive(train) bool
	notifyAboutDeparture()
}

// concrete mediator - stationManager
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
