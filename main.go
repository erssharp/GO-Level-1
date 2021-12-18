package main

import (
	"fmt"
	"math"
)

const (
	CalcRectangleSquareOption = 1
	CalcCircleDAndLOption     = 2
	ShowDigitsOption          = 3
)

func main() {
	fmt.Println("Choose an option:" +
		"\n1 - Calculate square of rectangle by 2 sides" +
		"\n2 - Calculate diameter and length of circle by it's square" +
		"\n3 - Show digits of 3-digit number")

	var option int
	if _, err := fmt.Scan(&option); err != nil {
		fmt.Println("Error:", err)
		return
	}

	switch option {
	case CalcRectangleSquareOption:
		var a, b int
		fmt.Print("Enter 2 sides od rectangle: ")
		if _, err := fmt.Scan(&a, &b); err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("\nS =", CalcRectangleSquare(a, b))
	case CalcCircleDAndLOption:
		var s float64
		fmt.Print("Enter square of circle: ")
		if _, err := fmt.Scan(&s); err != nil {
			fmt.Println("Error:", err)
			return
		}
		d, l := CalcCircleDAndL(s)
		fmt.Println("\nD =", d, "L =", l)
	case ShowDigitsOption:
		var value int
		fmt.Print("Enter value: ")
		if _, err := fmt.Scan(&value); err != nil {
			fmt.Println("Error:", err)
			return
		}
		first, second, third := ShowDigits(value)
		fmt.Println("\n", first, second, third)
	default:
		fmt.Println("You have entered wrong option")
	}
}

func ShowDigits(value int) (first int, second int, third int) {
	first = value / 100
	second = value / 10 % 10
	third = value % 10
	return
}

func CalcCircleDAndL(s float64) (d float64, l float64) {
	d = math.Sqrt(s/math.Pi) * 2
	l = d * math.Pi
	return
}

func CalcRectangleSquare(a, b int) int {
	return a * b
}
