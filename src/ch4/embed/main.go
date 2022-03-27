package main

import "fmt"


type Point struct {
	X int
	Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

var w Wheel
var w1 Wheel

func main() {
	w = Wheel{Circle{Point{8, 8}, 5}, 20}

	w1 = Wheel{
		Circle: Circle{
			Point: Point{
				X: 8,
				Y: 8,
			},
			Radius: 5,
		},
		Spokes: 20,
	}

	// # means that a synatx printer similar to go language
	// Can not include two anonymous members which are the same, it will cause name conflict
	fmt.Printf("%#v\n", w)
	w1.X = 42
	fmt.Printf("%v\n", w1)
}