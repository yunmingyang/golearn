// For avoiding circular import
package exportfortest_test

import (
	"fmt"
	"golearn/ch11/exportfortest"
	"testing"
)

var test123 = 0

func Test(t *testing.T) {
	exportfortest.IsR()
	saved := test123
	defer func() {
		test123 = saved
	}()

	test123 = 10
}

func Test1(t *testing.T) {
	fmt.Println(test123)
}
