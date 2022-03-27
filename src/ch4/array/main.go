package main

import (
	"crypto/sha256"
	"fmt"
)


func main() {
	var a [3]int
	fmt.Println(a[0])
	fmt.Println(a[len(a) - 1])

	for i, v := range a {
		fmt.Printf("%d %10d\n", i, v)
	}

	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

	// var q [3]int = [3]int{1, 2, 3}
	var r [3]int = [3]int{1, 2}
	fmt.Printf("%d\n", r[2])

	// Compute array length according to the number of initial value
	q := [...]int{1, 2, 3, 4}
	fmt.Printf("%T\n", q)

	m := [...]int{10: -1, -2, -3}
	fmt.Println(m)

	sha256Test()

	bArr := [32]byte{31: 1}
	fmt.Println(bArr)
	zero(&bArr)
	fmt.Println(bArr)

	months := [...]string{1: "January", 2: "Febuary", 3: "March", 4: "April", 12: "December"}
	m1 := months[1:3]
	for i := 0; i < len(months); i++ {
		fmt.Printf("%v\n", &months[i])
	}

	fmt.Printf("%v\n", "--------")

	for i := 0; i < len(m1); i++ {
		fmt.Printf("%v\n", &m1[i])
	}
	m2 := m1[:4]
	fmt.Println(m1)
	fmt.Println(m2)

	// slice of string will create a new string
	// & can not use for a[i] directly
	s := "Hello World"
	as := s[0: 4]
	fmt.Printf("%T %T\n", s, as)

	for i := 0; i < len(s); i++ {
		t := s[i]
		fmt.Printf("%T %v\n", t, &t)
	}

	for i := 0; i < len(as); i++ {
		t := as[i]
		fmt.Printf("%T  %v\n", t, &t)
	}

	// slice of string can not extend since its a new string
	// tt := as[:6]
	// fmt.Printf("%s\n", tt)

	// slice of []byte will also create a new []byte
	 bt := [3]byte{1, 2, 3}
	 bt1 := bt[0: 1]
	 fmt.Printf("%T %T %[1]v %[2]v\n", bt, bt1)

	 for i := 0; i < len(bt); i++ {
		t := bt[i]
		fmt.Printf("%T %v\n", t, &t)
	 }

	 for i := 0; i < len(bt1); i++ {
		t := bt1[i]
		fmt.Printf("%T %v\n", t, &t)
	 }

	// []byte could extend
	 bt2 := bt1[:2]
	 fmt.Printf("%T %v\n", bt2, bt2)
}


func sha256Test() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))

	fmt.Printf("%x\n%[1]b\n%x\n%[2]b\n%t\n%T\n", c1, c2, c1 == c2, c1)
	//exercise 4.1
	fmt.Printf("different bits: %d\n", diffByBytes(c1, c2))
}

func zero (ptr *[32]byte) {
	*ptr = [32]byte{}
}

func diffByBytes(c1, c2 [32]byte) int{
	count := 0

	for i := 0; i < 32; i ++ {
		count += diffByBits(c1[i], c2[i])
	}

	return count
}

func diffByBits(b1, b2 byte) int {
	count := 0

	for i := 0; i < 8; i ++ {
		mask := uint8(1) << i
		if b1 & mask != b2 & mask {
			count ++
		}
	}

	return count
}