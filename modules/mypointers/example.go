package mypointers

import "fmt"

func Example_ptr() {
	x := 42
	change_value(x)
	fmt.Println("Value:", x)
	change_ptr(&x)
	fmt.Println("Value Pointer:", x)

	fullname := "Kent"
	first_name(fullname)
	fmt.Println("New Name:", fullname)
	last_name(&fullname)
	fmt.Println("New Name Pointer:", fullname)
}

func change_value(val int) {
	val = 8
}

func change_ptr(ptr *int) {
	*ptr = 8
}

func first_name(name string) {
	name = "Clark"
}

func last_name(name_ptr *string) {
	*name_ptr = "Clark"
}
