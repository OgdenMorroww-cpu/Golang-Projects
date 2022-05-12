package mymaps

import "fmt"

func Go_maps() {

	comapny := map[string]int{
		"Toyota":      1932,
		"Lamborghini": 1963,
		"Ford":        1910,
		"Google":      1998,
		"Facebook":    2004,
	}
	for name, year := range comapny {
		fmt.Println("Company Name ->", name, "->", "Year Started", year)
	}
	nums := map[int]int{
		8: 12,
		2: 6,
		4: 9,
		5: 3,
	}
	delete(nums, 2)
	fmt.Println(nums[4] - nums[5])

}
