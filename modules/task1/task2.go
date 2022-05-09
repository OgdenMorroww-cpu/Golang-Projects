package task1

import "fmt"

func Get_planet_age_difference() {
	var age_on_earth int
	fmt.Scanln(&age_on_earth)

	mars := mars_age(age_on_earth * 365 / 687)
	fmt.Println(mars)
}

func mars_age(age_on_earth int) int {
	return age_on_earth
}
