package mypointers

import "fmt"

func Get_pointers() {
	x := 42
	y := &x
	*y = 45
	fmt.Println(*y)
	fmt.Println(x)

	first_name := "shedrack"
	last_name := &first_name
	*last_name = "Taylor"
	fmt.Println("First Name Address:", *last_name)
	fmt.Println(first_name)
	random(45)
	bank_name("Access Bank PLC")
	calaculate_sums(4)

}

func bank_name(bank_name_1 string) {
	bank_name_2 := &bank_name_1
	*bank_name_2 = "Zenith Bank PLC"
	fmt.Println("Bank name 2 address:", bank_name_2)
	fmt.Println("Bank name 2 value:", *bank_name_2)
	fmt.Println("Bank name 1 value:", bank_name_1)
}

func random(my_rand int) {
	my_rand = 45
	my_rand_2 := &my_rand
	*my_rand_2 = 100
	fmt.Println("My rand 2 address:", my_rand_2)
	fmt.Println("My rand 2 value:", *my_rand_2)
	fmt.Println("My rand:", my_rand)
}

func calaculate_sums(a int) {
	p := &a
	a += 2
	*p = *p - 1
	fmt.Println("The value of A:", a)
}
