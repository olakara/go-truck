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
	
	if err := truck.LoadCargo(); err != nil {
		return fmt.Errorf("error loading cargo: %w", err)
	}

	if err := truck.UnloadCargo(); err != nil {
		return fmt.Errorf("error unloading cargo: %w", err)

	}
	return nil
}

func main() {

	normalTruck := NormalTruck{id: "Normal Truck 1"}
	eletricTruck := EletricTruck{id: "Electric Truck 1", battery: 75}

	normalTruck.FuelUp(50)
	eletricTruck.Charge(25)

	if err := processTruck(&normalTruck); err != nil {

		log.Fatalf("Error processing truck: %v", err)
	}

	if err := processTruck(&eletricTruck); err != nil {
		log.Fatalf("Error processing truck: %v", err)
	}

	log.Println("Normal truck fuel: ", normalTruck.fuel)
	log.Println("Electric truck battery: ", eletricTruck.battery)	

	// person := make(map[string]any,0)
	// person["name"] = "John Doe"
	// person["age"] = 30

	// log.Println("Person: ", person)
	// age, exists := person["age"].(int)
	// if !exists {
	// 	log.Println("Age not found or not an int")
	// } else {
	// 	log.Println("Person age: ", age)
	// }

}
