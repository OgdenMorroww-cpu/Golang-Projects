package variadics

import "fmt"

func Variadics_fun() {
	sum(40, 60, 80, 100, 102, 204)
	multiple_names("Mark", "Musk", "Gate", "Jobs", "Bezos")
	cars := []string{"Bugatti", "Hennesy", "Koenisegg", "Aston-Martin"}
	multiple_names(cars...)
	v := []int{8, 5, 3}
	f(v...)
	s := []int{1, 2, 4, 6, 8}
	s[2] = s[1]
	s[3] = s[2] + s[0]
	fmt.Println(s[4] % s[3])
}

func sum(nums ...int) {
	total := 0
	for _, v := range nums {
		total += v
	}
	fmt.Println(total)
}

func multiple_names(names ...string) {
	fmt.Println("Names:", names)
}

func f(v ...int) {
	res := 20
	for _, a := range v {
		res -= a
	}
	fmt.Println(res)
}
