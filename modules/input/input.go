package input

import "fmt"

func Get_input() {
	var password string
	fmt.Scanln(&password)

	fmt.Println("Your Password is:", password)

	var input int
	fmt.Scanln(&input)

	fmt.Println(input * 10)
}
