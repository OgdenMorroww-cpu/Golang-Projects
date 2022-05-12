package arrays

import "fmt"

func Array_slice() {
	cars := [4]string{"bentley", "rolls-royce", "bmw", "hyundai"}
	var new_cars []string = cars[1:3]
	fmt.Println("Third Car:", new_cars)

	a := [5]int{1, 5, 2, 1, 8}
	var s []int = a[2:4]
	fmt.Println(s[1])

	users := make([]string, 4)
	users = append(users, "Mark", "Gayle", "Jace", "Kyle")
	fmt.Println(users)

	x := make([]int, 2)
	x = append(x, 3, 6, 2)
	fmt.Println(x[4])
}
