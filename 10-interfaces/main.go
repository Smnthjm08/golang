package main

import (
	"errors"
	"fmt"
	"log"
)

type Truck interface {
	LoadCargo() error
	UnLoadCargo() error
	ID() string
}

type NormalTruck struct {
	id    string
	cargo int
}

type ElectricTruck struct {
	id      string
	cargo   int
	battery float64
}

var (
	ErrNotImplemented = errors.New("not implemented")
	ErrTruckNotFound  = errors.New("truck not found")
)

func (t *NormalTruck) LoadCargo() error {
	return nil
}
func (t *NormalTruck) UnLoadCargo() error {
	return nil
}
func (t *NormalTruck) ID() string {
	return t.id
}

func (e *ElectricTruck) LoadCargo() error {
	return nil
}
func (e *ElectricTruck) UnLoadCargo() error {
	return nil
}
func (e *ElectricTruck) ID() string {
	return e.id
}

func processTruck(truck Truck) error {
	fmt.Printf("processing truck: %s\n", truck.ID())
	return nil
}

func main() {
	err := processTruck(&NormalTruck{id: "1"})
	if err != nil {
		log.Fatalf("error processing truck: %s", err)
	}

	err = processTruck(&ElectricTruck{id: "2"})
	if err != nil {
		log.Fatalf("error processing truck: %s", err)
	}
}
