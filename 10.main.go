package main

import (
	"fmt"
)

func calcArea(r float64) float64 {
	return 3.14 * r * r
}

func calcPermimeter(r float64) float64 {
	return 2 * 3.14 * r
}

func calcDiameter(r float64) float64 {
	return 2 * r
}

func main() {
	var query int
	var radius float64

	fmt.Print("Enter the radius of the circle: ")
	fmt.Scanf("%f ", &radius)
	fmt.Printf("Enter \n 1 = area \n 2 = perimeter \n 3 = diameter : ")
	fmt.Scanf("%d", &query)

	if query == 1 {
		fmt.Println("Results: ", calcArea(radius))
	} else if query == 2 {
		fmt.Println("Results: ", calcPermimeter(radius))
	} else if query == 3 {
		fmt.Println("Results: ", calcDiameter(radius))
	} else {
		fmt.Println("Invalid query")
	}
	fmt.Println("Thank you!")
}
