package main

import (
	"fmt"
	"math"

	"github.com/adarsh-jaiss/microservice-toll-calculator/types"
)

// beacuse it's an interface , we'll use er in the last of naming convention(nothing fancy)
type CalculatorServicer interface{
	CalculateDistance(types.OBUData) (float64,error)
}

type CalculatorService struct{
	points[][] float64

}

func NewCalculatorService() CalculatorServicer {
	return &CalculatorService{
		points: make([][]float64, 0),
	}

}

func(s *CalculatorService) CalculateDistance(data types.OBUData) (float64,error) {
	fmt.Println("calculating the distance")
	// distance := calculateDistance(data.Latitiude,data.Longitude)
	distance := 0.0
	if len(s.points)> 0{
		prevPoint := s.points[len(s.points)-1]
		distance = calculateDistance(prevPoint[0],prevPoint[1],data.Latitiude,data.Longitude)
	}
	s.points = append(s.points, []float64{data.Latitiude,data.Longitude})
	return distance,nil
}

func calculateDistance(x1,x2,y1,y2 float64) (float64) {
	return math.Sqrt(math.Pow(x2 - x1 ,2) + math.Pow(y2- y1 , 2))
}