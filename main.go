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
			fmt.Println("Error processing truck:", err)
			log.Fatalf("Error processing truck: %v", err)
		}

	}
}
