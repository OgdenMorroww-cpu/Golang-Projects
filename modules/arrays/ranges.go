package arrays

import "fmt"

func Get_range() {
	a := make([]int, 5)
	a[1] = 2
	a[2] = 3

	for i, v := range a {
		fmt.Println("Index:", i, "Value:", v)
	}

	hyper_cars := make([]string, 4)
	hyper_cars[0] = "Koenisegg"
	hyper_cars[1] = "Ferrari"
	hyper_cars[2] = "Lycan Hypersport"
	hyper_cars[3] = "Bugatti"

	for car_indexs, cars_values := range hyper_cars {
		fmt.Println("Index:", car_indexs, "Value:", cars_values)
	}

	todos := make([]string, 6)
	todos[0] = "Wash the cloth"
	todos[1] = "Walk my dog"
	todos[2] = "Book my flight"
	todos[3] = "Check Nnegi my babe"
	todos[4] = "Code my downloader projects"
	todos[5] = "Check my cat"

	for i, task := range todos {
		fmt.Println("Your Task No", i, "Is:", task)
	}
}
