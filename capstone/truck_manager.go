package main

import (
	"errors"
)

var ErrTruckNotFound = errors.New("truck not found")
var ErrTruckAlreadyExists = errors.New("truck already exists")

type FleetManager interface {
	AddTruck(id string, cargo int) error
	GetTruck(id string) (*Truck, error)
	RemoveTruck(id string) error
	UpdateTruckCargo(id string, cargo int) error
}

type Truck struct {
	ID    string
	Cargo int
}

type truckManager struct {
	trucks map[string]*Truck
}

func NewTruckManager() truckManager {
	return truckManager{
		trucks: make(map[string]*Truck),
	}
}

func (tm *truckManager) AddTruck(id string, cargo int) error {
	if _, exists := tm.trucks[id]; exists {
		return ErrTruckAlreadyExists
	}

	tm.trucks[id] = &Truck{ID: id, Cargo: cargo}
	return nil
}

func (tm *truckManager) GetTruck(id string) (*Truck, error) {
	if truck, exists := tm.trucks[id]; exists {
		return truck, nil
	}
	return nil, ErrTruckNotFound
}

func (tm *truckManager) RemoveTruck(id string) error {
	if _, exists := tm.trucks[id]; !exists {
		return ErrTruckNotFound
	}
	delete(tm.trucks, id)
	return nil
}
