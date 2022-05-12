package arrays

import "fmt"

func Mini_projects() {
	fmt.Println("Name?..")
	var get_fullName string
	fmt.Scanln(&get_fullName)

	for _, names := range get_fullName {
		fmt.Printf("%c \n", names)
	}

	result := 0
	nums := [3]int{2, 4, 6}
	for i, v := range nums {
		if i%2 == 0 {
			result += v
		}
	}
	fmt.Println(result)

}
