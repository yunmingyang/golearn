package main

import "fmt"


func main() {
	months := [...]string{1: "January",
						  2: "Febuary",
						  3: "March",
						  4: "April",
						  5: "May",
						  6: "June",
						  7: "July",
						  8: "August",
						  9: "September",
						  10: "October",
						  11: "November",
						  12: "December"}

	fmt.Println(months)

	Q2 := months[4:7]
	summer := months[6:9]

	fmt.Println(Q2)
	fmt.Println(summer)

	for _, s := range summer {
		for _, q := range Q2 {
			if q == s {
				fmt.Printf("%s appears in both\n", q)
			}
		}
	}

	// Out of range
	// fmt.Println(summer[:20])

	// Extend the slice
	endlessSummer := summer[:5]
	fmt.Println(endlessSummer)
}