package main

import (
	"errors"
	"sync"
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
	mu sync.RWMutex
}

func NewTruckManager() truckManager {
	return truckManager{
		trucks: make(map[string]*Truck),
	}
}

func (tm *truckManager) AddTruck(id string, cargo int) error {
	tm.mu.Lock()
	defer tm.mu.Unlock()
	if _, exists := tm.trucks[id]; exists {
		return ErrTruckAlreadyExists
	}

	tm.trucks[id] = &Truck{ID: id, Cargo: cargo}
	return nil
}

func (tm *truckManager) GetTruck(id string) (*Truck, error) {
	tm.mu.RLock()
	defer tm.mu.RUnlock()
	if truck, exists := tm.trucks[id]; exists {
		return truck, nil
	}
	return nil, ErrTruckNotFound
}

func (tm *truckManager) RemoveTruck(id string) error {
	tm.mu.Lock()
	defer tm.mu.Unlock()
	if _, exists := tm.trucks[id]; !exists {
		return ErrTruckNotFound
	}
	delete(tm.trucks, id)
	return nil
}

func (tm *truckManager) UpdateTruckCargo(id string, cargo int) error {
	tm.mu.Lock()
	defer tm.mu.Unlock()
	if truck, exists := tm.trucks[id]; exists {
		truck.Cargo = cargo
		return nil
	}
	return ErrTruckNotFound
}

