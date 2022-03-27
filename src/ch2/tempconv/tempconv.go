package tempconv

import (
	"fmt"
	"os"
	"strconv"
)


type Celsius float64
type Fahrenheit float64


const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC Celsius = 0
	BoilingC Celsius = 100
)


func (c Celsius) String() string { return fmt.Sprintf("%g%cC", c, '\u00B0') }


func (f Fahrenheit) String() string { return fmt.Sprintf("%g%cF", f, '\u00B0') }


func Tempconv() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)

		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}

		f := Fahrenheit(t)
		c := Celsius(t)

		fmt.Printf(
			"%s = %s, %s = %s\n",
			f,
			FtoC(f),
			c,
			CtoF(c),
		)
	}
}