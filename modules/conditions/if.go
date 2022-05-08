package conditions

import "fmt"

func Condition() {
	fmt.Println("Your Age....")
	var x int
	fmt.Scanln(&x)

	if x > 18 {
		fmt.Println("allowed")
	} else if x < 18 {
		fmt.Println("not allowed")
	}

	fmt.Println("Check Marital Status")
	var marital_status bool
	fmt.Scanln(&marital_status)

	if marital_status {
		fmt.Println("Shedrack is married to benita")
	} else {
		fmt.Println("Not valid")
	}

	if gpa := 6.8; gpa > 4.8 {
		fmt.Println("GPA is good")
	} else {
		fmt.Println("GPA is poor")
	}

	if temp := 8; temp/2 != 0 {
		fmt.Println(temp - 2)
	} else {
		fmt.Println(temp)
	}
}
