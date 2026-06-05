package main

import (
	"errors"
	"fmt"
	"log"
)

var (
	ErrTruckNotFound  = errors.New("NormalTruck not found")
	ErrNotImplemented = errors.New("Not implemented yet!")
)

type Truck interface {
	LoadCargo() error
	UnloadCargo() error
}

type NormalTruck struct {
	id string
}

func (t *NormalTruck) LoadCargo() error {
	return ErrTruckNotFound
}

func (t *NormalTruck) UnloadCargo() error {
	return nil
}

func processTruck(truck NormalTruck) error {
	fmt.Println("Processing truck: ", truck.id)

	if err := truck.LoadCargo(); err != nil {
		return fmt.Errorf("Error loading cargo: %w", err)
	}

	return ErrNotImplemented
}

func main() {
	trucks := []NormalTruck{
		{id: "NormalTruck-1"},
		{id: "NormalTruck-2"},
		{id: "NormalTruck-3"},
	}

	for _, truck := range trucks {
		fmt.Printf("NormalTruck %s arrived. \n", truck.id)

		if err := processTruck(truck); err != nil {
			log.Fatalf("Error processing truck: %s", err)
		}
	}
}
