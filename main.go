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

//processTruck handles the loading and unloading of a truck

func processTruck(truck Truck) error {
	//fmt.Println("Processing truck with ID:", truck.id)
	//return nil
	return ErrNotImplemented
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
			if errors.Is(err, ErrNotImplemented) {
				fmt.Println("Error: Functionality not implemented yet")
			}
			if errors.Is(err, ErrTruckNotFound) {
				fmt.Println("Error: Truck not found")
			}
			log.Fatalf("Error processing truck: %v", err)
		}

	}
}
