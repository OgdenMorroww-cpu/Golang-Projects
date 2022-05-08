package operators

import "fmt"

func Relations() {
	x := 42
	y := 8

	//Equal to
	fmt.Println(x == y)

	//Not equal to
	fmt.Println(x != y)

	//Greater than
	fmt.Println(x > y)

	//Less than
	fmt.Println(x < y)

	a := 50
	b := 25

	//Logical AND operator
	fmt.Println(a != b && a <= b)

	//Logical OR operator
	fmt.Println(a != b || a <= b)

	//Logical NOT operator
	fmt.Println(!(a > b))

	age := 15
	res := (age > 18) || age == 0
	fmt.Println("Result", res)
}
