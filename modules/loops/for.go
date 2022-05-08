package loops

import "fmt"

func Get_loops() {
	for index := 0; index < 5; index++ {
		fmt.Println(index)
	}

	sum := 1
	res := 0

	for sum <= 1000 {
		res += sum
		sum++
	}

	fmt.Println(res)

	big_num := 7
	small_num := 0

	switch {
	case big_num > 8:
		small_num++
	case big_num > 3 && big_num < 10:
		small_num++
	case big_num > 5:
		small_num++
	}
	fmt.Println(small_num)

	result := 0
	for i := 0; i < 10; i += 3 {
		result += i
	}
	fmt.Println(result)
}
