package parkinglot

import (
	"fmt"
	"sync"
)

type ParkingLot struct {
	levels []*Level
}

var (
	instance *ParkingLot
	once     sync.Once
)

func GetParkingLotInstance() *ParkingLot {
	once.Do(func() {
		fmt.Println("Creating ParkingLot instance")
		instance = &ParkingLot{levels: []*Level{}}
	})
	return instance
}

func (p *ParkingLot) AddLevel(level *Level) {
	p.levels = append(p.levels, level)
}

func (p *ParkingLot) ParkVehicle(vehicle Vehicle) bool {
	for _, level := range p.levels {
		if level.ParkVehicle(vehicle) {
			return true
		}
	}
	return false
}

func (p *ParkingLot) UnparkVehicle(vehicle Vehicle) bool {
	for _, level := range p.levels {
		if level.UnparkVehicle(vehicle) {
			return true
		}
	}
	return false
}

func (p *ParkingLot) DisplayAvailability() {
	for _, level := range p.levels {
		level.DisplayAvailability()
	}
}
