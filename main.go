package main

import (
	"errors"
	"fmt"
	"log"
	"time"
)

type Truck interface {
	LoadCargo() error
	UnloadCargo() error
}

type FuelTruck interface {
	FuelUp() error
}

type ChargeTruck interface {
	Charge() error
}

type NormalTruck struct {
	id    string
	cargo int
	fuel int
}

type EletricTruck struct {
	id    string
	cargo int
	battery int
}

var (
	ErrNotImplemented = errors.New("not implemented")
	ErrTruckNotFound  = errors.New("truck not found")
)

func (t *NormalTruck) LoadCargo() error {
	t.cargo++
	return nil
}

func (t *NormalTruck) UnloadCargo() error {
	t.cargo = 0
	return nil
}

func (t *NormalTruck) FuelUp(fuel int) error {
	if fuel <= 0 {
		return fmt.Errorf("invalid fuel amount: %d", fuel)
	}
	t.fuel += fuel
	log.Printf("Truck %s fueled up with %d liters\n", t.id, fuel)
	return nil
}


func (t *EletricTruck) LoadCargo() error {
	t.cargo++
	return nil
}

func (t *EletricTruck) UnloadCargo() error {
	t.cargo = 0
	return nil
}

func (t *EletricTruck) Charge(battery int) error {
	if battery <= 0 {
		return fmt.Errorf("invalid battery amount: %d", battery)
	}
	t.battery += battery
	log.Printf("Truck %s charged with %d%%\n", t.id, t.battery)
	return nil
}

//processTruck handles the loading and unloading of a truck

func processTruck(truck Truck) error {

	fmt.Printf("Processing truck: %+v\n", truck)

	// Simulate some proocessing time
	time.Sleep(time.Second)
	
	if err := truck.LoadCargo(); err != nil {
		return fmt.Errorf("error loading cargo: %w", err)
	}

	if err := truck.UnloadCargo(); err != nil {
		return fmt.Errorf("error unloading cargo: %w", err)

	}
	return nil
}

func processFleet(fleet []Truck) error {
	for _, truck := range fleet {
		processTruck(truck)
	}
	return nil
}

func main() {

	fleet := []Truck{
		&NormalTruck{id: "normal_truck_1", cargo: 0, fuel: 100},
		&EletricTruck{id: "electric_truck_1", cargo: 0, battery: 100},
		&NormalTruck{id: "normal_truck_2", cargo: 0, fuel: 50},
		&EletricTruck{id: "electric_truck_2", cargo: 0, battery: 80},
	}

	if err := processFleet(fleet); err != nil {
		log.Fatalf("Error processing fleet: %v", err)
	}

	fmt.Println("Fleet processed successfully")
}
