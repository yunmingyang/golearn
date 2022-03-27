package main

import "fmt"


type Celsius float64
type Fahrenheit float64


const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC Celsius = 0
	BoilingC Celsius = 100
)

// type conversion
func CtoF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }


func (c Celsius) String() string { return fmt.Sprintf("%g%cC", c, '\u00B0')}


func FtoC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9)}


func main() {
	var c Celsius

	fmt.Println(c.String())

	c = FtoC(212.0)
	fmt.Println(c.String())
	fmt.Printf("%v\n", c)
	fmt.Printf("%s\n", c)
	fmt.Println(c)
	fmt.Printf("%g\n", c)
	fmt.Println(float64(c))
}