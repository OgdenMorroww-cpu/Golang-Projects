package operators

import "fmt"

func ArithmeticOps() {
	x := 42
	y := 8

	//Addition
	res := x + y
	fmt.Println("Result:", res)

	//Subtraction
	res = x - y
	fmt.Println("Result:", res)

	//Division
	res = x / y
	fmt.Println("Result:", res)

	//Modulus remainder
	res = x % y
	fmt.Println("Results:", res)

	firstName := "Shedrack"
	lastName := "Abel"

	fullName := firstName + lastName
	fmt.Println("My full name is:", fullName)

	a := 8
	b := 3
	fmt.Println(a % b)

	first := 6
	second := first - 1
	first += second
	first %= second
	fmt.Println(first)
}
