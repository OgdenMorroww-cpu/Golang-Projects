package functions

import "fmt"

func Test_func() {
	for i := 4; i > 0; i-- {
		defer double_func(i)

	}
}

func double_func(a int) {
	fmt.Println(a * 2)
}
