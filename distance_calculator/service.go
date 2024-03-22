package main

import (
	"fmt"
	"math"

	"github.com/adarsh-jaiss/microservice-toll-calculator/types"
)

type CalculatorServicer interface {
	CalculateDistance(data types.OBUData) (float64, error)
}

type CalculatorService struct {
	Point [][]float64
}

func NewCalculatorService() CalculatorServicer {
	return &CalculatorService{
		Point: make([][]float64, 0),
	}

}

func (s *CalculatorService) CalculateDistance(data types.OBUData) (float64, error) {
	fmt.Println("calculating the distance")
	distance := 0.0
	if len(s.Point) > 0 {
		prevPoint := s.Point[len(s.Point)-1]
		distance = CalculateDistance(prevPoint[0], data.Latitiude, prevPoint[1], data.Longitude)
	}
	s.Point = append(s.Point, []float64{data.Latitiude, data.Longitude})
	return distance, nil
}

func CalculateDistance(x1, x2, y1, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}
