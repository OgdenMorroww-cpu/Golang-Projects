package functions

import "fmt"

func My_func() {
	fmt.Println("Hello func")
	num1, num2 := calculate(5)
	fmt.Println(num1 + num2)
	welcome("Hello world go functions")
	birth_day(1998, "Abel Shedrack Nicholas")
	result := get_sum(12, 45)
	fmt.Println("The sum is:", result)
	result1, reresult2, reresult3 := swap_func(20, 450, "Neil Mark")
	fmt.Printf("%v, %T\n", result1, result1)
	fmt.Printf("%v, %T\n", reresult2, reresult2)
	fmt.Printf("%v, %T\n", reresult3, reresult3)
	defer get_post(20, "Brad Traversy", true)
	fmt.Println("Delaying brad post today")
	fmt.Println("Start")
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("End")

}

func welcome(message string) string {
	fmt.Println("Your message is:", message)
	return message
}

func birth_day(birth_date int, celebrant string) (int, string) {
	current_age := 2022 - birth_date
	fmt.Println("You are", current_age, "years old")
	fmt.Println("Celebrant name is", celebrant)
	return birth_date, celebrant
}

func get_sum(x, y int) int {
	return x*y + 10
}

func swap_func(a, b int, null string) (int, int, string) {
	return a, b, null
}

func get_post(post_date int, post_author string, online_status bool) (int, string, bool) {
	fmt.Println("Post Date is:\n", post_date)
	fmt.Println("Post Author is:\n", post_author)
	fmt.Println("Online Status is:\n", online_status)
	return post_date, post_author, online_status
}

func calculate(a int) (int, int) {
	return a + 2, a + 1
}
