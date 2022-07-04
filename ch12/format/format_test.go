package format

import (
	"fmt"
	"testing"
	"time"
)



func Test(t *testing.T) {
	var x int64 = 1
	var d time.Duration = 1 * time.Nanosecond

	fmt.Println(Any(x))
	fmt.Println(Any(d))
	fmt.Println(Any([]int64{x}))
	fmt.Println(Any([]time.Duration{d}))
}