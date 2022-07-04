package main

import "fmt"



func main(){
	const freezingF, boilingF = 32.0, 212.0

	// degree centigrade unit uses unicode
	fmt.Printf("%g%cF = %g%cC\n", freezingF, '\u00B0', fToC(freezingF), '\u00B0')
	fmt.Printf("%g%cF = %g%cC\n", boilingF, '\u00B0', fToC(boilingF), '\u00B0')
}

func fToC(f float64) float64{
	return (f - 32) * 5 / 9
}
