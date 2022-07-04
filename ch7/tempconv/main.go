package main

import (
	"flag"
	"fmt"
)


type Celsius float64

func (c Celsius) String() string { return fmt.Sprintf("%g%cC", c, '\u00B0')}

type Fahrenheit float64

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

type celsiusFlag struct{ Celsius }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit{
	case "C", "\u00B0C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "\u00B0F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

var temp = CelsiusFlag("temp", 20.0, "the temperature")



func main() {
	flag.Parse()
	fmt.Println(*temp)
}