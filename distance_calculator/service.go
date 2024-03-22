package main

import (
	"fmt"

	"github.com/adarsh-jaiss/microservice-toll-calculator/types"
)

type CalculatorServicer interface {
	CalculateDistance(data types.OBUData) (float64, error)
}

type CalculatorService struct{}

func NewCalculatorService() CalculatorServicer {
	return &CalculatorService{}

}

func (s *CalculatorService) CalculateDistance(data types.OBUData) (float64, error) {
	fmt.Println("calculating the distance")
	return 0.0, nil

}
