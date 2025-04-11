package main

import (
	"errors"
	"fmt"
	"log"
)

type Truck struct {
	id string
}

var (
	ErrNotImplemented = errors.New("not implemented")
	ErrTruckNotFound  = errors.New("truck not found")
)

func (t Truck) UnloadCargo() error {
	return nil
}

func (t *Truck) LoadCargo() error {
	return nil
}

//processTruck handles the loading and unloading of a truck

func processTruck(truck Truck) error {

	fmt.Println("Processing truck:", truck.id)

	if err := truck.LoadCargo(); err != nil {
		return fmt.Errorf("error loading cargo: %w", err)
	}

	if err := truck.UnloadCargo(); err != nil {
		return fmt.Errorf("error unloading cargo: %w", err)

	}
	return nil
}

func main() {
	trucks := []Truck{
		{id: "T1"},
		{id: "T2"},
		{id: "T3"},
	}

	for _, truck := range trucks {
		fmt.Printf("Truck %s arrived\n", truck.id)
		if err := processTruck(truck); err != nil {

			log.Fatalf("Error processing truck: %v", err)
		}

	}
}
