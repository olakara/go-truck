package main

import "fmt"

type Truck struct {
	id string
}

//processTruck handles the loading and unloading of a truck

func processTruck(truck Truck) {
	fmt.Println("Processing truck with ID:", truck.id)
}

func main() {
	trucks := []Truck{
		{id: "T1"},
		{id: "T2"},
		{id: "T3"},
	}

	for _, truck := range trucks {
		fmt.Printf("Truck %s arrived\n", truck.id)
		processTruck(truck)
	}
}
