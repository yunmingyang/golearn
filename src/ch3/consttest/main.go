package main

import (
	"fmt"
	"net"
	"time"
)



func main(){
	const (
		noDelay time.Duration = 0
		timeout = 5 * time.Minute
	)
	fmt.Printf("%T %[1]v\n", noDelay)
	fmt.Printf("%T %[1]v\n", timeout)
	fmt.Printf("%T %[1]v\n", time.Minute)

	// When batch declaration, except the first one, the others initialization expressions could be omitted, and use the value and type of the previous one
	const (
		a = 1
		b
		c = 2
		d
	)
	fmt.Println(a, b, c, d)

	// // A example from package net
	// // iota const generator
	// type Flags uint
	// const (
	// 	FlagUp Flags = 1 << iota
	// 	FlagBroadcast
	// 	FlagLoopback
	// 	FlagPointToPoint
	// 	FlagMulticast
	// )

	var v net.Flags = net.FlagMulticast | net.FlagUp
	fmt.Printf("%b %t\n", v, isUp(v))
	turnDown(&v)
	fmt.Printf("%b %t\n", v, isUp(v))
	setBroadcast(&v)
	fmt.Printf("%b %t\n", v, isUp(v))
	fmt.Printf("%b %t\n", v, isCast(v))

	const (
		_ = 1 << (10 * iota)
		KiB
		MiB
		GiB
		TiB
		PiB
		EiB
		ZiB
		YiB
	)
	fmt.Printf("%b %b %b %b %b %b\n", KiB, MiB, GiB, TiB, PiB, EiB)

	const (
		_ = 1 << (10 * iota)
		KB
		MB
		GB
		TB
		PB
		EB
		ZB
		YB
	)
	fmt.Printf("%b %b %b %b %b %b\n", KB, MB, GB, TB, PB, EB)

	// compute when compiling
	fmt.Println(YiB/ZiB)

	const F float64 = 212
	fmt.Printf("%T\n", F)
	fmt.Printf("%T %[1]f\n", (F - 32) * 5 / 9)
	// fmt.Printf("%T %[1]f\n", 5 / 9 * (F - 32))
	fmt.Printf("%T %[1]f\n", 5.0 / 9.0 * (F - 32))
}

func isUp(f net.Flags) bool { return f & net.FlagUp == net.FlagUp }
func turnDown(f *net.Flags) { *f &^= net.FlagUp}
func setBroadcast(f *net.Flags) { *f |= net.FlagBroadcast}
func isCast(f net.Flags) bool { return f & (net.FlagBroadcast | net.FlagMulticast) != 0}