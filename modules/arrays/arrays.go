package arrays

import "fmt"

func Array_type() {
	a := [4]int{1, 2, 3, 4}
	fmt.Println("Arrays:", a)

	var odds [5]int
	odds[0] = 1
	odds[1] = 2
	odds[2] = 3
	odds[3] = 4
	odds[4] = 5

	fmt.Println("Odds:", odds)
}
