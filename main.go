package main

import (
	"errors"
	"fmt"
	"log"
)

type Truck interface {
	LoadCargo() error
	UnloadCargo() error
}

type NormalTruck struct {
	id    string
	cargo int
}

type EletricTruck struct {
	id    string
	cargo int
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

func (t *EletricTruck) LoadCargo() error {
	t.cargo++
	return nil
}

func (t *EletricTruck) UnloadCargo() error {
	t.cargo = 0
	return nil
}

//processTruck handles the loading and unloading of a truck

func processTruck(truck Truck) error {

	fmt.Printf("Processing truck: %v\n", truck)

	if err := truck.LoadCargo(); err != nil {
		return fmt.Errorf("error loading cargo: %w", err)
	}

	if err := truck.UnloadCargo(); err != nil {
		return fmt.Errorf("error unloading cargo: %w", err)

	}
	return nil
}

func main() {

	if err := processTruck(&NormalTruck{id: "Normal Truck 1"}); err != nil {

		log.Fatalf("Error processing truck: %v", err)
	}

	if err := processTruck(&EletricTruck{id: "Electric Truck 1"}); err != nil {
		log.Fatalf("Error processing truck: %v", err)
	}

}
