package main

import "fmt"


const boilingF = 212.0


func main(){
	var f = boilingF
	var c = (boilingF - 32) * 5 / 9

	fmt.Printf("boiling point: %g%cF or %g%cC\n", f, '\u00B0', c, '\u00B0')
}