package main

import (
	"fmt"
	"sort"
)


var m = make(map[string]int)


func main() {
	// Empty map
	ages := make(map[string]int)
	fmt.Println(ages)

	ages["alice"] = 31
	ages["charlie"] = 42
	fmt.Println(ages)

	ages1 := map[string]int{
		"alice": 1,
		"charlie": 2,
	}
	fmt.Println(ages1)

	// Empty map
	ages2 := map[string]int{}
	fmt.Println(ages2)

	fmt.Println(ages["alice"])

	delete(ages, "alice")
	fmt.Println(ages)

	// Safe operation, since failed finding will retrun corresponding zero value
	ages["bob"] = ages["bob"] + 1
	fmt.Println(ages)

	ages["k"] += 1
	ages["s"]++
	fmt.Println(ages)

	// Error since map element not a variable
	// _ = &ages["bob"]

	for name, age := range ages {
		fmt.Printf("%s is %d years old\n", name, age)
	}

	// var names []string
	names := make([]string, 0, len(ages))
	for name := range ages {
		names = append(names, name)
	}

	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}

	var ages3 map[string]int
	fmt.Println(ages3 == nil)
	fmt.Println(len(ages3) == 0)
	// panic
	// ages3["carlo"] = 21

	age, ok := ages["Bob"]
	if !ok {
		fmt.Printf("no such key: %s\t%d\n", "Bob", age)
	}

	ages["sss"] = 0
	age, ok = ages["sss"]
	if ok {
		fmt.Printf("%s exist, but its value is %d\n", "sss", age)
	}

	age, ok = ages["sss1"]
	if !ok {
		fmt.Printf("%s doesn't exist, and its value is %d\n", "sss1", age)
	}

	// Reslove the problem which is the ok may be only used in the if block
	if age, ok := ages["Slince"]; !ok {
		fmt.Printf("no such key: %s\t%d\n", "Slince", age)
	}

	fmt.Println(equal(map[string]int{}, map[string]int{"t": 1}))
}


func Add (list []string) {m[k(list)]++}

func Count (list []string) int {return m[k(list)]}

func k (list []string) string {return fmt.Sprintf("%q", list)}

func equal (x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}

	for xkey, xval := range x {
		if yval, ok := y[xkey]; !ok || yval != xval {
			return false
		}
	}

	return true
}