package main

import "fmt"



type Point struct { X,Y float64 }

func (p Point) Add(q Point) Point { return Point{p.X + q.X, p.Y + q.Y}}
func (p Point) Sub(q Point) Point { return Point{p.X - q.X, p.Y - q.Y}}


type Path []Point

func (path Path) TranslateBy(offset Point, add bool) {
	var op func(p, q Point) Point

	if add {
		op = Point.Add
	} else {
		op = Point.Sub
	}

	for i := range path {
		path[i] = op(path[i], offset)
	}
}

func main() {
	p := Path{
		Point{1, 2},
		Point{3, 4},
		Point{5, 6},
	}

	p.TranslateBy(Point{10, 10}, true)

	fmt.Println("path: ", p)

	fmt.Println("s", 1 << 0)
	fmt.Println("s", 1 << 1)
	a := uint64(0)
	fmt.Printf("binary pring: %b\n", a)
	fmt.Println(a)
	a |= 1 << 0
	fmt.Printf("binary pring: %b\n", a)
	a |= 1 << 1
	fmt.Printf("binary pring: %b\n", a)
}