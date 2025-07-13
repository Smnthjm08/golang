package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

type contextKey string

var UserIDKey contextKey = "userID"

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

func (t *NormalTruck) LoadCargo() error   { return nil }
func (t *NormalTruck) UnLoadCargo() error { return nil }
func (t *NormalTruck) ID() string         { return t.id }

func (e *ElectricTruck) LoadCargo() error   { return nil }
func (e *ElectricTruck) UnLoadCargo() error { return nil }
func (e *ElectricTruck) ID() string         { return e.id }

func processTruck(ctx context.Context, truck Truck) error {
	// userID := ctx.Value(UserIDKey)
	userID := ctx.Value(UserIDKey)
	// log.Println(userID)

	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()

	fmt.Printf("User %v processing truck: %s\n", userID, truck.ID())
	return nil
}

func processFleet(ctx context.Context, fleet []Truck) error {
	for _, truck := range fleet {
		if err := processTruck(ctx, truck); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, UserIDKey, 42)

	fleet := []Truck{
		&NormalTruck{id: "NT1", cargo: 0},
		&ElectricTruck{id: "ET1", cargo: 0, battery: 100},
		&NormalTruck{id: "NT@", cargo: 0},
		&ElectricTruck{id: "NT2", cargo: 0, battery: 100},
		&NormalTruck{id: "NT3", cargo: 0},
	}

	if err := processFleet(ctx, fleet); err != nil {
		fmt.Printf("Error processing fleet: %v\n", err)
		return
	}

	// if err := processTruck(ctx, &NormalTruck{id: "1"}); err != nil {
	// 	log.Fatalf("error processing truck: %s", err)
	// }

	// if err := processTruck(ctx, &ElectricTruck{id: "2"}); err != nil {
	// 	log.Fatalf("error processing truck: %s", err)
	// }
}
