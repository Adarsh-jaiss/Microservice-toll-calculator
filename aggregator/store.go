package main

import "github.com/adarsh-jaiss/microservice-toll-calculator/types"

type MemoryStore struct {}

func(m *MemoryStore) Insert(d types.Distance) error {
	return nil
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{}	
	
}