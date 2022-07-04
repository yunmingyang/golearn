package main

import (
	"fmt"
	"time"
)


type Employee struct {
	ID int
	Name string
	Address string
	DoB time.Time
	Postion string
	Salary int
	ManagerID int
}

type Point struct {
	x, y int
}

type address struct {
	hostname string
	port int
}

var dilbert Employee

func Scale(p Point, factor int) Point {
	return Point{p.x * factor, p.y * factor}
}


func main() {
	dilbert.Salary -= 5000
	fmt.Println(dilbert.Salary)

	postion := &dilbert.Postion
	*postion = "Senior" + dilbert.Postion
	fmt.Println(*postion)

	var employeeOfTheMonth *Employee = &dilbert
	// The same with "(*employeeOfTheMonth).Postion += "(private team player)""
	employeeOfTheMonth.Postion += "(private team player)"
	fmt.Println(dilbert.Postion)

	p := Point{1, 2}
	fmt.Println(p)
	p1 := Point{x: 3, y: 4}
	fmt.Println(p1)
	pp := &Point{5, 6}
	// Struct pointer
	fmt.Printf("%T %[1]v\n", pp)

	fmt.Println(Scale(Point{1, 2}, 5))

	p2 := Point{1, 2}
	q2 := Point{2, 3}
	fmt.Println(p2.x == q2.x && p2.y == q2.y)
	fmt.Println(p2 == q2)

	hits := make(map[address]int)
	hits[address{"golang.org", 443}]++
	fmt.Println(hits)
}