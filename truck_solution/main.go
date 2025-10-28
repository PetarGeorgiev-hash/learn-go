package main

import (
	"errors"
	"sync"
)

var ErrTruckNotFound = errors.New("truck not found")

type FleetManager interface {
	AddTruck(id string, cargo int) error
	GetTruck(id string) (Truck, error)
	RemoveTruck(id string) error
	UpdateTruckCargo(id string, cargo int) error
}

type Truck struct {
	ID    string
	Cargo int
}

type truckManager struct {
	trucks map[string]*Truck
	sync.RWMutex
}

func (t *truckManager) AddTruck(id string, cargo int) error {
	t.Lock()
	defer t.Unlock()

	if _, exists := t.trucks[id]; exists {
		return errors.New("truck with this ID: %s already exists")
	}
	t.trucks[id] = &Truck{ID: id, Cargo: cargo}
	return nil
}

func (t *truckManager) GetTruck(id string) (Truck, error) {
	t.RLock()
	defer t.RUnlock()

	if truck, exists := t.trucks[id]; exists {
		return *truck, nil
	}

	return Truck{}, ErrTruckNotFound
}

func (t *truckManager) RemoveTruck(id string) error {
	t.Lock()
	defer t.Unlock()
	if _, exists := t.trucks[id]; !exists {
		return ErrTruckNotFound
	}
	delete(t.trucks, id)
	return nil
}

func (t *truckManager) UpdateTruckCargo(id string, cargo int) error {
	t.Lock()
	defer t.Unlock()
	if truck, exists := t.trucks[id]; exists {
		truck.Cargo = cargo
		return nil
	}
	return ErrTruckNotFound
}

func NewTruckManager() truckManager {
	return truckManager{
		trucks: make(map[string]*Truck),
	}
}
