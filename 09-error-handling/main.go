package main

import (
	"errors"
	"fmt"
	"log"
)

type Truck interface {
	LoadCargo() error
	UnLoadCargo() error
}

type NormalTruck struct {
	id    string
	cargo int
}

var (
	ErrNotImplemented = errors.New("not implemented")
	ErrTruckNotFound  = errors.New("truck not found")
)

type ElectricTruck struct {
	id      string
	cargo   int
	battery float64
}

func (t *NormalTruck) LoadCargo() error {
	return nil
}

func (t *NormalTruck) UnLoadCargo() error {
	return nil
}

// Processing error for any function is good handling
func processTruck(truck NormalTruck) error {
	// fmt.Printf("processing truck: %s\n", truck.id)

	// if err := truck.LoadCargo(); err != nil {
	// 	return fmt.Errorf("error loading cargo: %w", err)
	// }

	return nil
	// return ErrNotImplemented
}

func main() {
	trucks := []NormalTruck{
		{id: "Truck-1"},
		{id: "Truck-2"},
		{id: "Truck-3"},
		{id: "Truck-4"},
	}

	// eTrucks := []ElectricTruck{
	// 	{id: "Electric-truck-1"},
	// }

	for _, truck := range trucks {
		fmt.Printf("Truck %s arrived.\n", truck.id)
		// err := processTruck(truck)

		// err := processTruck(truck)
		// if err != nil {
		// 	log.Fatalf("Error processing truck: %s", err)
		// }

		if err := processTruck(truck); err != nil {
			if errors.Is(err, ErrNotImplemented) {
				// we do this
				log.Fatal("TRUE")
			}
			log.Fatalf("Error processing truck: %s", err)
		}
	}

}
