package task1

import "fmt"

func MyTask() {

	for num_range := 1; num_range <= 3; num_range++ {
		var numbers int
		fmt.Scanln(&numbers)

		switch numbers {
		case 0:
			fmt.Println("Zero")
		case 1:
			fmt.Println("One")
		case 2:
			fmt.Println("Two")
		case 3:
			fmt.Println("Three")
		case 4:
			fmt.Println("Four")
		case 5:
			fmt.Println("Five")
		case 6:
			fmt.Println("Six")
		case 7:
			fmt.Println("Seven")
		case 8:
			fmt.Println("Eight")
		case 9:
			fmt.Println("Nine")
		default:
			fmt.Println("Ten")
		}
	}

}
