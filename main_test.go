package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	t.Run("processTruck", func(t *testing.T) {
		
		t.Run("Should load and unload a normal truck", func(t *testing.T) {
			truck := &NormalTruck{id: "normal_truck_1", cargo: 0, fuel: 100}
			if err := processTruck(truck); err != nil {
				t.Errorf("Failed to process normal truck: %v", err)
			}

			// assert that the truck's cargo is 0 after unloading
			if truck.cargo != 0 {
				t.Fatalf("Expected cargo to be 0 after unloading, got %d", truck.cargo)
			}

		})
	})
}