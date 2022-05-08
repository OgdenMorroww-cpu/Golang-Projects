package conditions

import "fmt"

func Get_switch() {
	fmt.Println("Get num value....")

	var num int
	fmt.Scanln(&num)

	switch num {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println(num)
	}

	x := 8

	switch y := x % 2; y {
	case 0:
		x -= 1
	case 1:
		x += 1
	}
	fmt.Println(x)
}
